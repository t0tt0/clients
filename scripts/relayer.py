import subprocess
import time

nt = 2

def rcmd():
	logfile = open('relayerlog', 'w')
	nonce = hex(12)
	ret1 = subprocess.Popen('node ./twoconnections.js -t ' + nonce + ' -c 0', shell=True, stdout=logfile, stderr=logfile, encoding="utf-8")
	while(1):
		if ret1.poll() == 0:
			print("checker done")
			ret1.terminate()
			break;
		else:
			time.sleep(1)
			
	

def poster():
	posterlogfile = open('posterlog', 'w')
	print("enter relayer.................")
	# checkerlogfile = open('checkerlog', 'w')
	for i in range(0, 1, 1):


		ret1 = subprocess.Popen('node ./twoconnections.js -t ' + hex(1) + ' -c 0', shell=True, stdout=posterlogfile, stderr=posterlogfile, encoding="utf-8")
		for i  in range(0, 8, 1):
			if ret1.poll() == 0:
				print("checker done")
				ret1.terminate()
				break;
			else:
				time.sleep(1)


		ret = subprocess.Popen("python3 /Users/taoyuechen/go/src/github.com/relayer/go-ves/cmd/ves-client/playbook.py", shell=True, stdout=posterlogfile,stderr=posterlogfile, encoding="utf-8")
		for x in range (0, 8, 1):
			if ret.poll() == 0:
				print("hhhhhhhhh success")
				ret.terminate()
				break;
			else:
				time.sleep(1)

		# time.sleep(10)

		nonce = hex(i)
		

def checker():
	
	time.sleep(10)
	for i in range(0, nt, 1):
		ret1 = subprocess.Popen('node ./twoconnections.js', shell=True, stdout=checkerlogfile, stderr=checkerlogfile, encoding="utf-8")
		while(1):
			if ret1.poll() == 0:
				print("checker done")
				ret1.terminate()
				break;
			else:
				time.sleep(1)

if __name__ == '__main__':


	# parser = argparse.ArgumentParser(description='Process some integers.')

	# parser.add_argument('--nonce', '-n', help='an integer for nonce', required = True)
	# parser.add_argument('--nonce1', '-x', help='an integer for nonce1', required = True)


	# args = parser.parse_args()

	# rcmd()
	poster()






