package proof

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"sync"

	log "github.com/inconshreveable/log15"
	"github.com/pgc/utils"
)

// Global params
var (
	Compressed    = 33
	uint64ByteLen = 8
	hashMapFile   = filepath.Join(os.Getenv("HOME"), "hashMap")
)

var point2Index = make(map[uint64]int, 0)

// BuildAndLoadMapIfNotExist tries to build map if it's not exist.
// Load map if exist
func BuildAndLoadMapIfNotExist(g *utils.ECPoint, rangeLen, tunning, roNum int) error {
	if _, err := os.Stat(hashMapFile); os.IsNotExist(err) {
		log.Info("build hash map, cost about 1 minute")
		return buildHashMap(g, rangeLen, tunning, roNum)
	}
	return loadHashMap(rangeLen, tunning)
}

func buildHashMap(g *utils.ECPoint, rangeLen, tunning, roNum int) error {
	giantStepSize := 2 << (rangeLen/2 + tunning - 1)

	if giantStepSize%roNum != 0 {
		return errors.New("Thread assignment fails")
	}

	l := giantStepSize / roNum
	buffer := make([]byte, giantStepSize*uint64ByteLen)

	wg := sync.WaitGroup{}
	wg.Add(roNum)

	for i := 0; i < roNum; i++ {
		go func(i int) {
			startPoint := new(utils.ECPoint).ScalarMult(g, big.NewInt(int64(i*l)))
			calPoints(startPoint, g, buffer, l, i*l)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// write to file
	if err := ioutil.WriteFile(hashMapFile, buffer, 0644); err != nil {
		return err
	}

	// load data
	return loadHashMap(rangeLen, tunning)
}

func LoadMap(rangeLen, tunning int) error {
	return loadHashMap(rangeLen, tunning)
}

func loadHashMap(rangeLen, tunning int) error {
	giantStepSize := 2 << (rangeLen/2 + tunning - 1)
	bytesLen := giantStepSize * uint64ByteLen

	// read file
	bytes, err := ioutil.ReadFile(hashMapFile)
	if err != nil {
		return err
	}

	if bytesLen != len(bytes) {
		return errors.New("Invalid hash map file bytes len")
	}

	for i := 0; i < giantStepSize; i++ {
		keyBz := bytes[i*uint64ByteLen : (i+1)*uint64ByteLen]
		key := binary.LittleEndian.Uint64(keyBz)
		point2Index[key] = i
	}

	return nil
}

// ShanksDlog decrypts point using shanks algorithm.
func ShanksDlog(g, msg *utils.ECPoint, rangeLen, tunning int) (*big.Int, error) {
	giantStepSize := 2 << (rangeLen/2 + tunning - 1)
	loop := 2 << (rangeLen/2 - tunning - 1)

	giantStep := new(utils.ECPoint).ScalarMult(g, big.NewInt(int64(giantStepSize)))
	giantStep.Negation(giantStep)

	dstPoint := msg.Copy()
	if giantStepSize != len(point2Index) {
		return nil, errors.New("Hash map isn't loaded")
	}

	i, j := 0, 0
	find := false
	for ; j < loop; j++ {
		msgKey := dstPoint.MurmurHashToUint64()
		r, ok := point2Index[msgKey]
		if ok {
			find = true
			i = r
			break
		}

		dstPoint.Add(dstPoint, giantStep)
	}

	if !find {
		return nil, errors.New("The DLOG is not found in the specified range")
	}

	return big.NewInt(int64(i + j*giantStepSize)), nil
}

func calPoints(start, g *utils.ECPoint, buffer []byte, l, startIndex int) {
	for i := 0; i < l; i++ {
		index := (startIndex + i) * uint64ByteLen
		bz := make([]byte, uint64ByteLen)
		binary.LittleEndian.PutUint64(bz, start.MurmurHashToUint64())
		copy(buffer[index:index+uint64ByteLen], bz)
		start.Add(start, g)
	}
}
