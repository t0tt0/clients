import numpy as np
import random
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
import subprocess
import time
import matplotlib.pyplot as plt

nt = 1
fl = np.linspace(1, nt, nt, dtype = int)
cf = np.linspace(1, 3, 3, dtype = int)

numcf = np.zeros((nt,3), dtype = int)
numother = np.zeros((nt, 3), dtype=int)
numtot = np.zeros((nt, 3), dtype=int)

pconfir = np.zeros((nt, 3), dtype = int)
pcrash = np.zeros((nt, 3), dtype = int)
cc = np.zeros((nt, 3), dtype = int)

# ttt = np.zeros((nt, 3), dtype = int)

fato = open("fpato.txt", "a+")
fsuc = open("fpsuc.txt", "a+")

pc = np.linspace(0, 0.8, 5, dtype = float)
pnlato_4pc = np.linspace(0, 0.8, 5, dtype = float)
pnunsuc_4pc = np.linspace(0, 0.8, 5, dtype = float)
ppnunsuc_4pc = np.linspace(0, 0.8, 5, dtype = float)



pnlato = np.linspace(0, 0.8, 5, dtype = float)
pnunsuc = np.linspace(0, 0.8, 5, dtype = float)

pnlato_hy = np.linspace(0, 0.8, 5, dtype = float)
pnunsuc_hy = np.linspace(0, 0.8, 5, dtype = float)

def selectTrans():
	

	# print(pc)
	for p in range (0, 100, 20):
		nlato = 0
		nunsuc = 0
		nnunsuc = 0

		po = float(p/100)

		nonce = 0
		nonce1 = 0

		ps = 1-(1-po)*(1-po)*(1-po)*(1-po)*(1-po)*(1-po)*(1-po)*(1-po)*(1-po)*(1-po)
		for i in range (0, 100, 1):
			ran = random.uniform(0, 1)
			if ran <= ps:
				ran1 = random.uniform(0, 1)

				senderlogfile = open('senderlog', 'w')			

				ret1 = subprocess.Popen("node ./toeth.js -t " + hex(1) + " -c 0", shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				if ran1 < 1-(1-po)*(1-po)*(1-po):
					ret1.terminate()
					nonce1 = nonce1 + 1
				else:
					while(1):
						if ret1.poll() == 0:
							print("success tother")
							ret1.terminate()
							break;
						else:
							time.sleep(1)
				#== 
					ret2 = subprocess.Popen("python3 /Users/taoyuechen/go/src/github.com/HyperService-Consortium1/go-ves/cmd/ves-client/playbook.py", shell=True, stdout=senderlogfile,stderr=senderlogfile, encoding="utf-8")
					ret3 = subprocess.Popen('node ./twoconnections.js -t ' + hex(1) + ' -c 0', shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
					# time.sleep(1)
					ret2.terminate()
					ret3.terminate()
					nonce = nonce + 1
				
			else:
				continue

		clogfile = open('clog', 'w')
		# print("python3 stopmining.py -n " + str(non), end=' ')
		r1 = subprocess.Popen("python3 stopmining.py -n " + str(nonce) + " -x " + str(nonce1), shell=True, stdout=clogfile, stderr=clogfile, encoding="utf-8")
		time.sleep(2)
		# r1.terminate()

	for p in range (0, 100, 20):
		nlato = 0
		nunsuc = 0

		po = float(p/100)

		nonce = 0
		nonce1 = 0
		ps = 1-(1-po)*(1-po)*(1-po)*(1-po)*(1-po)*(1-po)
		for i in range (0, 100, 1):
			senderlogfile = open('senderlog', 'w')
			ran = random.uniform(0, 1)
			if ran <= ps:
				ran1 = random.uniform(0, 1)

				ret1 = subprocess.Popen("node ./toeth.js -t " + hex(nonce) + " -c 0", shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				ret1.terminate()
				ret2 = subprocess.Popen("python3 relayer.py", shell=True)
				nonce1 = nonce1 + 1

		clogfile = open('clog', 'w')
		# print("python3 stopmining.py -n " + str(non), end=' ')
		r1 = subprocess.Popen("python3 stopmining.py -n " + str(nonce) + " -x " + str(nonce1), shell=True, stdout=clogfile, stderr=clogfile, encoding="utf-8")
		time.sleep(2)
		r1.terminate()

	for p in range (0, 100, 20):
		nlato = 0
		nunsuc = 0

		po = float(p/100)

		nonce = 0
		nonce1 = 0

		ps = 1-(1-po)*(1-po)*(1-po)*(1-po)*(1-po)*(1-po)*(1-po)*(1-po)
		for i in range (0, 100, 1):
			ran = random.uniform(0, 1)

			if ran <= ps:
				ran1 = random.uniform(0, 1)

				if ran1 < 0.5:
					ret = subprocess.Popen("python3 /Users/taoyuechen/go/src/github.com/HyperService-Consortium/go-ves/cmd/ves-client/playbook.py", shell=True, stdout=senderlogfile,stderr=senderlogfile, encoding="utf-8")
					ret.terminate()
					nonce1 = nonce1 + 1
				else:
					ret = subprocess.Popen("python3 /Users/taoyuechen/go/src/github.com/HyperService-Consortium/go-ves/cmd/ves-client/playbook.py", shell=True, stdout=senderlogfile,stderr=senderlogfile, encoding="utf-8")
					time.sleep(3)
					ret.terminate()
					nonce = nonce + 1

		clogfile = open('clog', 'w')
		# print("python3 stopmining.py -n " + str(non), end=' ')
		r1 = subprocess.Popen("python3 stopmining_hy.py -n " + str(nonce) + " -t " + nonce1, shell=True, stdout=clogfile, stderr=clogfile, encoding="utf-8")
		time.sleep(2)
		r1.terminate()

		pnlato_hy[int(p/20)] = float(nlato/100)
		pnunsuc_hy[int(p/20)] = float(nunsuc/100)

	







if __name__ == '__main__':
	selectTrans()

