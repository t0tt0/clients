import numpy as np
import random
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
import subprocess
import time
from matplotlib.ticker import FuncFormatter
# %config InlineBackend.figure_formats
# %matplotlib inline
# %config InlineBackend.print_figure_kwargs = {'bbox_inches':None}

#plt.rcParams["figure.figsize"] = (20,5)

nt = 100
fl = np.linspace(1, int(3*nt), 100, dtype = int)
cf = np.linspace(30, 100, 3, dtype = float)

#print(cf)

cf = [x/100 for x in cf]

numcf = np.zeros((nt,3), dtype = int)
numother = np.zeros((nt, 3), dtype=int)
numtot = np.zeros((nt, 3), dtype=int)

pconfir = np.zeros((nt, 3), dtype = int)
pcrash = np.zeros((nt, 3), dtype = int)
cc = np.zeros((nt, 3), dtype = int)

# ttt = np.zeros((nt, 3), dtype = int)

fato = open("fato.txt", "a+")
fsuc = open("fsuc.txt", "a+")

def changex(temp, position):
	return int(temp/50)


def selectTrans():
	
	for i in range (1, nt+1, 1):
		fa = i*3
		for j in range (1, 4, 1):
			confir = j * i
			other = fa - confir
			o1 = other
			if confir >= nt:
				confir = nt
			if other >= 2*nt:
				other = 2*nt
			if other >= nt:
				o1 = nt
			resultlist = random.sample(range(0, nt), confir)
			resultlist1 = random.sample(range(0, nt), o1)
			intersection = list(set(resultlist) | set(resultlist1))
			numcf[i-1][j-1] = len(resultlist)
			numother[i-1][j-1] = len(intersection) - len(resultlist)
			numtot[i-1][j-1] = len(intersection)
			
			pconfir[i-1][j-1] = len(list(set(resultlist).difference(set(resultlist1))))
			pcrash[i-1][j-1] = len(list(set(resultlist1).difference(set(resultlist))))
			# pcrash[i-1][j-1] = numother[i-1][j-1]
			cc[i-1][j-1] = len(list(set(resultlist).intersection(set(resultlist1))))


	

	for i in range (1, nt+1, 1):
		for j in range (1, 4, 1):
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
				ret2 = subprocess.Popen("python3 relayer.py", shell=True)
				for x in range (0, 50, 1):
					if ret2.poll() == 0:
						print("success playbook")
						ret2.terminate()
						break;
					else:
						time.sleep(1)
				time.sleep(1)
				ret1.terminate()
				# nonce = hex(i)
				# ret1 = subprocess.Popen('node ./twoconnections.js -t ' + nonce + ' -c 0', shell=True, stdout=posterlogfile, stderr=posterlogfile, encoding="utf-8")

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
				nonce = hex(i)
				ret1 = subprocess.Popen("node ./toeth.js -t " + nonce + " -c 0", shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				while(1):
					if ret1.poll() == 0:
						print("success tother")
						ret1.terminate()
						break;
					else:
						time.sleep(1)
				# ret11 = subprocess.Popen("cd /Users/taoyuechen/go/src/github.com/HyperService-Consortium/go-ves/cmd/ves-client", shell=True)
				# time.sleep(1)
				# ret11.terminate()
				ret2 = subprocess.Popen("python3 relayer.py", shell=True)
				# ddl = datetime.now().Add(time.Millisecond * 800)
				for x in range (0, 50, 1):
					if ret2.poll() == 0:
						print("success playbook")
						ret2.terminate()
						break;
					else:
						time.sleep(1)
				# while(1):
				# 	if ret1.poll() == 0:
				# 		print("success playbook")
				# 		ret1.terminate()
				# 		break;
				# 	else:
				# 		time.sleep(1)
				# time.sleep(3)
			for i in range (0, ncc, 1):
				nonce = hex(i)
				ret1 = subprocess.Popen("node ./toeth.js -t " + nonce + " -c 0", shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
				while(1):
					if ret1.poll() == 0:
						print("success tother")
						ret1.terminate()
						break;
					else:
						time.sleep(1)
				ret2 = subprocess.Popen("python3 relayer.py", shell=True)
				for x in range (0, 50, 1):
					if ret2.poll() == 0:
						print("success playbook")
						ret2.terminate()
						break;
					else:
						time.sleep(1)
				time.sleep(3)
				ret1.terminate()
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
			time.sleep(5) #shall be altered

			non = npcrash + npconfir + ncc
			clogfile = open('clog', 'w')
			print("python3 stopmining.py -n " + str(non), end=' ')
			r1 = subprocess.Popen("python3 stopmining.py -n " + str(non), shell=True, stdout=clogfile, stderr=clogfile, encoding="utf-8")
			time.sleep(2)
			r1.terminate()

		fato.write("\n")
		fsuc.write("\n")


	fato.close()
	fsuc.close()

	# # print(numother)
	# print(pcrash)
	# print(pconfir)
	# print(cc)

	# print(numother)
	# print()

	
	numother1 = np.transpose(numother)
	numtot1 = np.transpose(numtot)

	fig = plt.figure(figsize=(20, 5))
	# ax1 = fig.add_subplot(121, projection='3d')
	# X, Y = np.meshgrid(fl, cf)

	# ax1.plot_surface(X, Y, numother1, rstride=1, cstride=1)
	# ax1.zaxis.set_tick_params( labelsize=20)
	# ax1.zaxis.set_ticks(np.arange(0, 100, 30))
	# ax1.xaxis.set_tick_params(labelsize=20)
	# ax1.xaxis.set_ticks(np.arange(0, 300, 50))
	# ax1.xaxis.set_major_formatter(FuncFormatter(changex))
	# ax1.yaxis.set_tick_params(labelsize=20)
	# ax1.yaxis.set_ticks(np.arange(0.3, 1.1, 0.3))
	# ax1.set_xlabel('F*100', fontsize = 20, labelpad = 21)
	# ax1.set_ylabel('$c$', fontsize = 20, labelpad = 21)
	# ax1.yaxis._axinfo['label']['space_factor'] = 3.0
	# ax1.set_zlabel('E', fontsize = 20, labelpad = 21, rotation = 0)

	# # plt.title('4pc', fontsize = 20)

	# ax1.set_title('4pc', fontsize=20)

        # plt.title('4pc', fontsize = 20)

	# ax2 = fig.gca(projection='3d')
	ax2 = fig.add_subplot(111, projection = '3d')
	# ax2 = fig.add_axes([0, 0, 1, 1], projection='3d')
	# ax2 = plt.axes(projection='3d')
	X, Y = np.meshgrid(fl, cf)
	ax2.plot_surface(X, Y, numtot1, rstride=1, cstride=1)


	ax2.zaxis.set_tick_params( labelsize=20)
	ax2.zaxis.set_ticks(np.arange(0, 100, 30))

	ax2.xaxis.set_tick_params(labelsize=20)
	ax2.xaxis.set_ticks(np.arange(0, 300, 50))
	ax2.xaxis.set_major_formatter(FuncFormatter(changex))

	ax2.yaxis.set_tick_params( labelsize=20)
	ax2.yaxis.set_ticks(np.arange(0.3, 1.1, 0.3))
	ax2.set_xlabel('F*100', fontsize = 20, labelpad = 15)
	ax2.set_ylabel('$c$', fontsize = 20, labelpad = 15)
	# ax2.yaxis._axinfo['label']['space_factor'] = 3.0
	ax2.set_zlabel('E', fontsize = 20, labelpad = 15, rotation = 0)
	# ax2.set_title('Hyperservice', fontsize = 20)
	plt.title('Hyperservice', fontsize = 20)

	# plt.show()
	plt.savefig('Hyperserviceand4pc.pdf', bbox_inches = 'tight')



	# plt.savefig('4pc.pdf', bbox_inches='tight')



if __name__ == '__main__':
	selectTrans()
