#!/usr/bin/python
# -*- coding: UTF-8 -*-

import random



def gran():
	sum = 0
	sum1 = 0
	for i in range (10):
		sum = sum + random.randint(15, 30)
		sum1 = sum1 + random.randint(30, 60)
		i = i + 1
	print(sum)
	print("\n")
	print(sum1)


if __name__ == "__main__":
    gran()

# print( random.randint(15,30) )        # 产生 1 到 10 的一个整数型随机数  
# print( random.random() )             # 产生 0 到 1 之间的随机浮点数
# print( random.uniform(1.1,5.4) )     # 产生  1.1 到 5.4 之间的随机浮点数，区间可以不是整数
# print( random.choice('tomorrow') )   # 从序列中随机选取一个元素
# print( random.randrange(1,100,2) )   # 生成从1到100的间隔为2的随机整数

# a=[1,3,5,6,7]                # 将序列a中的元素顺序打乱
# random.shuffle(a)
# print(a)
