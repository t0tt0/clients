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

fato = open("fpato_con.txt", "a+")
fsuc = open("fpsuc_con.txt", "a+")

pc = np.linspace(0, 0.8, 5, dtype = float)
pnlato_4pc = np.linspace(0, 0.8, 5, dtype = float)
pnunsuc_4pc = np.linspace(0, 0.8, 5, dtype = float)

pnlato = np.linspace(0, 0.8, 5, dtype = float)
pnunsuc = np.linspace(0, 0.8, 5, dtype = float)

pnlato_hy = np.linspace(0, 0.8, 5, dtype = float)
pnunsuc_hy = np.linspace(0, 0.8, 5, dtype = float)

def selectTrans():
	

	# print(pc)
	for p in range (0, 100, 20):
		nlato = 0
		nunsuc = 0

		po = float(p/100)

		nonce = 0

		ps = 1-(1-po)*(1-po)
		for i in range (0, 100, 1):
			ran = random.uniform(0, 1)
			if ran <= ps:
				ran1 = random.uniform(0, 1)

				senderlogfile = open('senderlog', 'w')

				stoplogfile = open('stoplog', 'w')
				r = subprocess.Popen("node ./stopmining.js", shell=True, stdout=stoplogfile, stderr=stoplogfile, encoding="utf-8")
				while(1):
					if r.poll() == 0:
						print("success stop mining")
						r.terminate()
						break;
					else:
						time.sleep(1)
				ret1 = subprocess.Popen("node ./toeth.js -t " + 1 + " -c 0", shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				while(1):
					if ret1.poll() == 0:
						print("success tother")
						ret1.terminate()
						break;
					else:
						time.sleep(1)
				ret3 = subprocess.Popen('node ./twoconnections.js -t ' + 1 + ' -c 0', shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				while(1):
					if ret3.poll() == 0:
						print("success commit")
						ret3.terminate()
						break;
					else:
						time.sleep(1)
				ret2 = subprocess.Popen("python3 /Users/taoyuechen/go/src/github.com/HyperService-Consortium1/go-ves/cmd/ves-client/playbook.py", shell=True, stdout=senderlogfile,stderr=senderlogfile, encoding="utf-8")
				for x in range (0, 8, 1):
					if ret2.poll() == 0:
						print("success playbook, and commit")
						ret2.terminate()
						break;
					else:
						time.sleep(1)

				nonce = nonce + 1

				startlogfile = open('startlog', 'w')
				r = subprocess.Popen("node ./startmining.js", shell=True, stdout=startlogfile, stderr=startlogfile, encoding="utf-8")
				while(1):
					if r.poll() == 0:
						print("success start mining")
						r.terminate()
						break;
					else:
						time.sleep(1)
				# time.sleep(5)#may change

				
			else:
				continue

		clogfile = open('clog', 'w')
		# print("python3 stopmining.py -n " + str(non), end=' ')
#new cal file
		r1 = subprocess.Popen("python3 stopmining.py -n " + str(nonce) + " -x " + str(0), shell=True, stdout=clogfile, stderr=clogfile, encoding="utf-8")
		time.sleep(2)
		r1.terminate()

# 	for p in range (0, 100, 20):
# 		nlato = 0
# 		nunsuc = 0

# 		po = float(p/100)

# 		nonce = 0
# 		ps = 1-(1-po)*(1-po)
# 		for i in range (0, 100, 1):
# 			senderlogfile = open('senderlog', 'w')

# 			ran = random.uniform(0, 1)
# 			if ran <= ps:
# 				# ran1 = random.uniform(0, 1)

# 				stoplogfile = open('stoplog', 'w')
# 				r = subprocess.Popen("node ./stopmining.js", shell=True, stdout=stoplogfile, stderr=stoplogfile, encoding="utf-8")
# 				while(1):
# 					if r.poll() == 0:
# 						print("success stop mining")
# 						r.terminate()
# 						break;
# 					else:
# 						time.sleep(1)

# 				ret1 = subprocess.Popen("node ./toeth.js -t " + nonce + " -c 0", shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
# 				while(1):
# 					if ret1.poll() == 0:
# 						print("success tother")
# 						ret1.terminate()
# 						break;
# 					else:
# 						time.sleep(1)
# 				# ret11 = subprocess.Popen("cd /Users/taoyuechen/go/src/github.com/HyperService-Consortium1/go-ves/cmd/ves-client", shell=True)
# 				# time.sleep(1)
# 				# ret11.terminate()
# 				ret2 = subprocess.Popen("python3 relayer.py", shell=True)
# 				# ddl = datetime.now().Add(time.Millisecond * 800)
# 				for x in range (0, 50, 1):
# 					if ret2.poll() == 0:
# 						print("success playbook")
# 						ret2.terminate()
# 						break;
# 					else:
# 						time.sleep(1)
# 				startlogfile = open('startlog', 'w')
# 				r = subprocess.Popen("node ./startmining.js", shell=True, stdout=startlogfile, stderr=startlogfile, encoding="utf-8")
# 				while(1):
# 					if r.poll() == 0:
# 						print("success start mining")
# 						r.terminate()
# 						break;
# 					else:
# 						time.sleep(1)
# 				nonce = nonce + 1


# 		clogfile = open('clog', 'w')
# 		# print("python3 stopmining.py -n " + str(non), end=' ')
# #new cal file......
# 		r1 = subprocess.Popen("python3 stopmining.py -n " + str(nonce) + " -c 0", shell=True, stdout=clogfile, stderr=clogfile, encoding="utf-8")
# 		time.sleep(2)
# 		r1.terminate()

# 	for p in range (0, 100, 20):
# 		nlato = 0
# 		nunsuc = 0

# 		po = float(p/100)

# 		nonce = 0
# 		nonce1 = 0

# 		ps = 1-(1-po)*(1-po)
# 		for i in range (0, 100, 1):
# 			ran = random.uniform(0, 1)

# 			if ran <= ps:
# 				stoplogfile = open('stoplog', 'w')
# 				r = subprocess.Popen("node ./stopmining.js", shell=True, stdout=stoplogfile, stderr=stoplogfile, encoding="utf-8")
# 				while(1):
# 					if r.poll() == 0:
# 						print("success stop mining")
# 						r.terminate()
# 						break;
# 					else:
# 						time.sleep(1)

# 				ret1 = subprocess.Popen("python3 /Users/taoyuechen/go/src/github.com/HyperService-Consortium1/go-ves/cmd/ves-client/playbook.py", shell=True, stdout=senderlogfile,stderr=senderlogfile, encoding="utf-8")
# 				# ddl = datetime.now().Add(time.Millisecond * 800)
# 				for x in range (0, 8, 1):
# 					if ret1.poll() == 0:
# 						print("success playbook")
# 						ret1.terminate()
# 						break;
# 					else:
# 						time.sleep(1)

# 				nonce = nonce + 1

# 				startlogfile = open('startlog', 'w')
# 				rrr1 = subprocess.Popen("node ./startmining.js", shell=True, stdout=startlogfile, stderr=startlogfile, encoding="utf-8")
# 				while(1):
# 					if rrr1.poll() == 0:
# 						print("success start mining")
# 						rrr1.terminate()
# 						break;
# 					else:
# 						time.sleep(1)

# 		clogfile = open('clog', 'w')
# 		# print("python3 stopmining.py -n " + str(non), end=' ')
# #new cal file
# 		r1 = subprocess.Popen("python3 stopmining_hy.py -n " + str(nonce), shell=True, stdout=clogfile, stderr=clogfile, encoding="utf-8")
# 		time.sleep(2)
# 		r1.terminate()






if __name__ == '__main__':
	selectTrans()

