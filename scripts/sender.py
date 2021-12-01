import subprocess
import time

nt = 2

def rcmd():
	logfile = open('relayerlog', 'w')
	ret = subprocess.Popen("ls -l", shell=True, stdout=logfile,stderr=logfile, encoding="utf-8")
	# ret.wait(2)
	while(1):
		if(ret.poll() == 0):
			print("success", ret.communicate())
			break
		else:
			ret.wait(1)



def sender():
	senderlogfile = open('senderlog', 'w')
	for i in range(0, nt, 1):
		nonce = hex(i)
		ret1 = subprocess.Popen("node ./toeth.js -t " + nonce + " -c 0", shell=True, stdout=senderlogfile, stderr=senderlogfile, encoding="utf-8")
		while(1):
			if ret1.poll() == 0:
				print("success tother")
				ret1.terminate()
				break;
			else:
				time.sleep(1)
		# ret = subprocess.Popen("python3 /Users/taoyuechen/go/src/github.com/HyperService-Consortium/go-ves/cmd/ves-client/playbook.py", shell=True, stdout=senderlogfile,stderr=senderlogfile, encoding="utf-8")
		# while(1):
		# 	if ret.poll() == 0:
		# 		print("success playbook")
		# 		ret.terminate()
		# 		break;
		# 	else:
		# 		time.sleep(1)

		time.sleep(3)

def checker():
	checkerlogfile = open('checkerlog', 'w')
	time.sleep(10)
	# for i in range(0, nt, 1):
	# 	ret1 = subprocess.Popen('node ./twoconnections.js', shell=True, stdout=checkerlogfile, stderr=checkerlogfile, encoding="utf-8")
	# 	while(1):
	# 		if ret1.poll() == 0:
	# 			print("checker done")
	# 			ret1.terminate()
	# 			break;
	# 		else:
	# 			ret1.wait(1)

if __name__ == '__main__':
	sender()