# 当前版本的bulletproof(basic)实现
// 依照实际代码
1. commitment V = g*r + h*m.; 其中g为生成key的基点, g,h 均随机生成.
2. 在bulletproof中:
* 生成proof时: 
 ** A = h*alpha + G*AL + H*AR;  S = h*rho + G*SL + H*SR
 ** T1 = g*(random) + h*t1; T2 = g*(random) + h*t2;
* 验证proof时: 
 right1 = V*(z^2) + h*dleta_yz + ...
 right2 = h*mu + G*L + hPrime*R.

 # 之前版本的代码中

 1. commitment V = g*m + h*r; h 作为生成key的基点.
 2. bulletproof中:
 * 生成proof: 
  ** A = h*alpha + G*AL + H*AR;  S = h*rho + G*SL + H*SR 
  ** T1 = g*t1 + h *(random); T2 = g*t2+h*(random);
 * 验证proof:
right = V*(z^2) + g*dleta_yz + ...

# bulletproof论文
1. commitment V = g*m + h*r.
2. 生成proof
A = h*alpha + G*AL + H*AR;  S = h*rho + G*SL + H*SR
T1 = g*t1 + h*r; T2 = g*t2 + h*r + ...
3. 验证proof时
right = V*(z^2) + g*dleta_yz + ...
right2 = h*mu + G*L + hPrime*R.

# 问题
1. 现在的代码与bulletproof的论文是否有不一致的地方? 
之前的代码与论文的的一致,V均等于g*m+h*r(h为生成key点);
其他地方代码也保持一致;在当前最新代码中, 对g,h点进行了调换, V=g*r+h*m.(g为生成key点), 但所有代码在生成A,S点及计算right2时使用的均为h点, 没有进行调换, 其他使用g,h的位置进行了调换.