import numpy as np
import random
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
import subprocess
import time

nt = 1 #nt = 1-100
fl = np.linspace(1, nt, nt, dtype = int)
cf = np.linspace(1, 3, 3, dtype = int)

numcf = np.zeros((nt,3), dtype = int)
numother = np.zeros((nt, 3), dtype=int)
numtot = np.zeros((nt, 3), dtype=int)

# pconfir = np.zeros((nt, 3), dtype = int)
# pcrash = np.zeros((nt, 3), dtype = int)
# cc = np.zeros((nt, 3), dtype = int)

# ttt = np.zeros((nt, 3), dtype = int)

fato = open("fato.txt", "a+")
fsuc = open("fsuc.txt", "a+")

def selectTrans():
	
# 	for i in range (1, nt+1, 1):
# 		fa = i*3
# 		for j in range (1, 4, 1):
# 			confir = j * i
# 			other = fa - confir
# 			o1 = other
# 			if confir >= nt:
# 				confir = nt
# 			if other >= 2*nt:
# 				other = 2*nt
# 			if other >= nt:
# 				o1 = nt
# 			resultlist = random.sample(range(0, nt), confir)
# 			resultlist1 = random.sample(range(0, nt), o1)
# 			intersection = list(set(resultlist) | set(resultlist1))
# 			numcf[i-1][j-1] = len(resultlist)
# 			numother[i-1][j-1] = len(intersection) - len(resultlist)
# 			numtot[i-1][j-1] = len(intersection)
			
# 			pconfir[i-1][j-1] = len(list(set(resultlist).difference(set(resultlist1))))
# 			pcrash[i-1][j-1] = len(list(set(resultlist1).difference(set(resultlist))))
# 			pcrash[i-1][j-1] = numother[i-1][j-1]
# 			cc[i-1][j-1] = len(list(set(resultlist).intersection(set(resultlist1))))

	pcrash = [ [ 2,1,0,], [ 4,2,0,], [ 6,3,0,], [ 8,4,0,], [ 10,4,0,], [ 11,6,0,], [ 12,6,0,], [ 15,7,0,], [ 16,5,0,], [ 18,10,0,], [ 16,8,0,], [ 24,10,0,], [ 20,11,0,], [ 25,12,0,], [ 26,10,0,], [ 29,11,0,], [ 30,11,0,], [ 29,10,0,], [ 31,13,0,], [ 33,11,0,], [ 31,12,0,], [ 40,12,0,], [ 38,10,0,], [ 36,14,0,], [ 36,14,0,], [ 41,12,0,], [ 43,16,0,], [ 37,13,0,], [ 42,10,0,], [ 45,14,0,], [ 45,10,0,], [ 41,12,0,], [ 47,12,0,], [ 45,13,0,], [ 43,9,0,], [ 45,7,0,], [ 48,10,0,], [ 52,7,0,], [ 46,7,0,], [ 48,9,0,], [ 48,6,0,], [ 46,5,0,], [ 50,4,0,], [ 46,5,0,], [ 49,4,0,], [ 50,5,0,], [ 49,5,0,], [ 50,3,0,], [ 50,1,0,], [ 50,0,0,], [ 49,0,0,], [ 48,0,0,], [ 47,0,0,], [ 46,0,0,], [ 45,0,0,], [ 44,0,0,], [ 43,0,0,], [ 42,0,0,], [ 41,0,0,], [ 40,0,0,], [ 39,0,0,], [ 38,0,0,], [ 37,0,0,], [ 36,0,0,], [ 35,0,0,], [ 34,0,0,], [ 33,0,0,], [ 32,0,0,], [ 31,0,0,], [ 30,0,0,], [ 29,0,0,], [ 28,0,0,], [ 27,0,0,], [ 26,0,0,], [ 25,0,0,], [ 24,0,0,], [ 23,0,0,], [ 22,0,0,], [ 21,0,0,], [ 20,0,0,], [ 19,0,0,], [ 18,0,0,], [ 17,0,0,], [ 16,0,0,], [ 15,0,0,], [ 14,0,0,], [ 13,0,0,], [ 12,0,0,], [ 11,0,0,], [ 10,0,0,], [ 9,0,0,], [ 8,0,0,], [ 7,0,0,], [ 6,0,0,], [ 5,0,0,], [ 4,0,0,], [ 3,0,0,], [ 2,0,0,], [ 1,0,0,], [ 0,0,0,]]

	pconfir = [ [ 1,2,3,], [ 2,4,6,], [ 3,6,9,], [ 4,8,12,], [ 5,9,15,], [ 5,12,18,], [ 5,13,21,], [ 7,15,24,], [ 7,14,27,], [ 8,20,30,], [ 5,19,33,], [ 12,22,36,], [ 7,24,39,], [ 11,26,42,], [ 11,25,45,], [ 13,27,48,], [ 13,28,51,], [ 11,28,54,], [ 12,32,57,], [ 13,31,60,], [ 10,33,63,], [ 18,34,66,], [ 15,33,69,], [ 12,38,72,], [ 11,39,75,], [ 15,38,78,], [ 16,43,81,], [ 9,41,84,], [ 13,39,87,], [ 15,44,90,], [ 14,41,93,], [ 9,44,96,], [ 14,45,99,], [ 11,47,100,], [ 8,44,100,], [ 9,43,100,], [ 11,47,100,], [ 14,45,100,], [ 7,46,100,], [ 8,49,100,], [ 7,47,100,], [ 4,47,100,], [ 7,47,100,], [ 2,49,100,], [ 4,49,100,], [ 4,51,100,], [ 2,52,100,], [ 2,51,100,], [ 1,50,100,], [ 0,50,100,], [ 0,49,100,], [ 0,48,100,], [ 0,47,100,], [ 0,46,100,], [ 0,45,100,], [ 0,44,100,], [ 0,43,100,], [ 0,42,100,], [ 0,41,100,], [ 0,40,100,], [ 0,39,100,], [ 0,38,100,], [ 0,37,100,], [ 0,36,100,], [ 0,35,100,], [ 0,34,100,], [ 0,33,100,], [ 0,32,100,], [ 0,31,100,], [ 0,30,100,], [ 0,29,100,], [ 0,28,100,], [ 0,27,100,], [ 0,26,100,], [ 0,25,100,], [ 0,24,100,], [ 0,23,100,], [ 0,22,100,], [ 0,21,100,], [ 0,20,100,], [ 0,19,100,], [ 0,18,100,], [ 0,17,100,], [ 0,16,100,], [ 0,15,100,], [ 0,14,100,], [ 0,13,100,], [ 0,12,100,], [ 0,11,100,], [ 0,10,100,], [ 0,9,100,], [ 0,8,100,], [ 0,7,100,], [ 0,6,100,], [ 0,5,100,], [ 0,4,100,], [ 0,3,100,], [ 0,2,100,], [ 0,1,100,], [ 0,0,100,]]
	cc = [ [ 0,0,0,], [ 0,0,0,], [ 0,0,0,], [ 0,0,0,], [ 0,1,0,], [ 1,0,0,], [ 2,1,0,], [ 1,1,0,], [ 2,4,0,], [ 2,0,0,], [ 6,3,0,], [ 0,2,0,], [ 6,2,0,], [ 3,2,0,], [ 4,5,0,], [ 3,5,0,], [ 4,6,0,], [ 7,8,0,], [ 7,6,0,], [ 7,9,0,], [ 11,9,0,], [ 4,10,0,], [ 8,13,0,], [ 12,10,0,], [ 14,11,0,], [ 11,14,0,], [ 11,11,0,], [ 19,15,0,], [ 16,19,0,], [ 15,16,0,], [ 17,21,0,], [ 23,20,0,], [ 19,21,0,], [ 23,21,0,], [ 27,26,0,], [ 27,29,0,], [ 26,27,0,], [ 24,31,0,], [ 32,32,0,], [ 32,31,0,], [ 34,35,0,], [ 38,37,0,], [ 36,39,0,], [ 42,39,0,], [ 41,41,0,], [ 42,41,0,], [ 45,42,0,], [ 46,45,0,], [ 48,48,0,], [ 50,50,0,], [ 51,51,0,], [ 52,52,0,], [ 53,53,0,], [ 54,54,0,], [ 55,55,0,], [ 56,56,0,], [ 57,57,0,], [ 58,58,0,], [ 59,59,0,], [ 60,60,0,], [ 61,61,0,], [ 62,62,0,], [ 63,63,0,], [ 64,64,0,], [ 65,65,0,], [ 66,66,0,], [ 67,67,0,], [ 68,68,0,], [ 69,69,0,], [ 70,70,0,], [ 71,71,0,], [ 72,72,0,], [ 73,73,0,], [ 74,74,0,], [ 75,75,0,], [ 76,76,0,], [ 77,77,0,], [ 78,78,0,], [ 79,79,0,], [ 80,80,0,], [ 81,81,0,], [ 82,82,0,], [ 83,83,0,], [ 84,84,0,], [ 85,85,0,], [ 86,86,0,], [ 87,87,0,], [ 88,88,0,], [ 89,89,0,], [ 90,90,0,], [ 91,91,0,], [ 92,92,0,], [ 93,93,0,], [ 94,94,0,], [ 95,95,0,], [ 96,96,0,], [ 97,97,0,], [ 98,98,0,], [ 99,99,0,], [ 100,100,0,]]

	for i in range (nt, nt+1, 1):
		for j in range (2, 3, 1): #(1, 2, 1) (2,3,1) (3, 4, 1)
			npcrash = pcrash[i-1][j-1]
			npconfir = pconfir[i-1][j-1]
			ncc = cc[i-1][j-1]

			# for k in range (1, npcrash, 1):
			senderlogfile = open('senderlog', 'w')			

			for i in range(0, npcrash, 1):
				nonce = hex(i)
				ret1 = subprocess.Popen("node ./toeth.js -t " + nonce + " -c 0", shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				while(1):
					if ret1.poll() == 0:
						print("success tother")
						ret1.terminate()
						break;
					else:
						time.sleep(1)
				#== 
				ret2 = subprocess.Popen("python3 /Users/taoyuechen/go/src/github.com/HyperService-Consortium1/go-ves/cmd/ves-client/playbook.py", shell=True, stdout=senderlogfile,stderr=senderlogfile, encoding="utf-8")
				ret3 = subprocess.Popen('node ./twoconnections.js -t ' + nonce + ' -c 0', shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				# time.sleep(1)
				ret2.terminate()
				ret3.terminate()
				
				time.sleep(3)
			
			stoplogfile = open('stoplog', 'w')
			r = subprocess.Popen("node ./stopmining.js", shell=True, stdout=stoplogfile, stderr=stoplogfile, encoding="utf-8")
			while(1):
				if r.poll() == 0:
					print("success stop mining")
					r.terminate()
					break;
				else:
					time.sleep(1)

			for i in range (0, npconfir, 1):
				nonce =  hex(i + npcrash)
				print("confir ti")
				ret1 = subprocess.Popen("node ./toeth.js -t " + nonce + " -c 0", shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				print("confir ti end")
				for x in range (0, 10, 1):
					if ret1.poll() == 0:
						print("success tother")
						ret1.terminate()
						break;
					else:
						time.sleep(1)
				print("quit sending in confir")
				ret3 = subprocess.Popen('node ./twoconnections.js -t ' + nonce + ' -c 0', shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				while(1):
					if ret3.poll() == 0:
						print("success commit")
						ret3.terminate()
						break;
					else:
						time.sleep(1)
				print("quit two connections...")
				#hyper1
				ret2 = subprocess.Popen("python3 /Users/taoyuechen/go/src/github.com/HyperService-Consortium1/go-ves/cmd/ves-client/playbook.py", shell=True, stdout=senderlogfile,stderr=senderlogfile, encoding="utf-8")
				for x in range (0, 8, 1):
					if ret2.poll() == 0:
						print("success playbook, and commit")
						ret2.terminate()
						break;
					else:
						time.sleep(1)
				time.sleep(3)


			for i in range (0, ncc, 1):
				print("in ncc..................")
				nonce =  hex(i + npcrash)
				#hyper1
				ret1 = subprocess.Popen("node ./toeth.js -t " + nonce + " -c 0", shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				while(1):
					if ret1.poll() == 0:
						print("success tother")
						ret1.terminate()
						break;
					else:
						time.sleep(1)
				#hyper1
				ret2 = subprocess.Popen("python3 /Users/taoyuechen/go/src/github.com/HyperService-Consortium1/go-ves/cmd/ves-client/playbook.py", shell=True, stdout=senderlogfile,stderr=senderlogfile, encoding="utf-8")
				ret2.terminate()
				# for x in range (0, 8, 1):
				# 	if ret2.poll() == 0:
				# 		print("success playbook, and commit")
				# 		ret2.terminate()
				# 		break;
				# 	else:
				# 		time.sleep(1)
				ret3 = subprocess.Popen('node ./twoconnections.js -t ' + nonce + ' -c 0', shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				ret3.terminate()
				# time.sleep(1)
				# ret2.terminate()
				
				# while(1):
				# 	if ret.poll() == 0:
				# 		print("success playbook")
				# 		ret.terminate()
				# 		break;
				# 	else:
				# 		time.sleep(1)
				time.sleep(3)

			startlogfile = open('startlog', 'w')
			r = subprocess.Popen("node ./startmining.js", shell=True, stdout=startlogfile, stderr=startlogfile, encoding="utf-8")
			while(1):
				if r.poll() == 0:
					print("success start mining")
					r.terminate()
					break;
				else:
					time.sleep(1)
			time.sleep(5)#may change

			non = npcrash + npconfir + ncc
			clogfile = open('clog', 'w')
			# print("python3 stopmining.py -n " + str(non), end=' ')
			r1 = subprocess.Popen("python3 stopmining.py -n " + str(non), shell=True, stdout=clogfile, stderr=clogfile, encoding="utf-8")
			while(1):
				if r1.poll() == 0:
					print("get success")
					r1.terminate()
					break;
				else:
					time.sleep(1)
			# r1.terminate()

		fato.write("\n")
		fsuc.write("\n")


	fato.close()
	fsuc.close()




if __name__ == '__main__':
	selectTrans()

