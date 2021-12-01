from web3 import Web3
import rlp
from eth_utils import keccak, to_checksum_address, to_bytes
import argparse

def stop1():
	web3 = Web3(Web3.HTTPProvider("http://127.0.0.1:24765"))
	print(web3.isConnected())
	print("....1")

	web3.geth.miner.start(1)

def calato(totalnonce):
	web3 = Web3(Web3.HTTPProvider("http://127.0.0.1:24765"))
	web31 = Web3(Web3.HTTPProvider("http://127.0.0.1:24768"))
	print(web3.isConnected())
	print(web31.isConnected())
	print("....1")


	sender = bytes.fromhex("94abee207bc81159831e3c6b3a6722d019e8e13e")
	sender1 = bytes.fromhex("6f55fcf81624e908da2ba3c83341ff707b652760")

	poster = bytes.fromhex("94abee207bc81159831e3c6b3a6722d019e8e13e")
	poster1 = bytes.fromhex("6f55fcf81624e908da2ba3c83341ff707b652760")

	ato = 0
	suc = 0

	fato = open("fato.txt", "w")
	fsuc = open("fsuc.txt", "w")

	for i in range(0, totalnonce, 1):
		raw = rlp.encode([sender, i])
		h = keccak(raw)
		address = h[12:]
		# print(to_checksum_address(address) )

		raw1 = rlp.encode([sender1, i])
		h1 = keccak(raw1)
		address1 = h1[12:]
		# print(to_checksum_address(address1) )

		v = web3.eth.get_balance(address)
		v1 = web31.eth.get_balance(address1)

		# u = web3.eth.get_transaction_count()

		if v > 0 and v1 > 0:
			suc = suc + 1
			ato = ato + 1
			if i == totalnonce -1:
				print(suc, end = ' ')
				print(ato, end = ' ')
				fato.write(str(ato) + " ")
				fsuc.write(str(suc) + " ")
		elif v == 0 and v1 == 0:
			ato = ato + 1
			if i == totalnonce -1:
				print(suc, end = ' ')
				print(ato, end = ' ')
				fato.write(str(ato) + " ")
				fsuc.write(str(suc) + " ")
		else:
			if i == totalnonce -1:
				print(suc, end = ' ')
				print(ato, end = ' ')
				fato.write(str(ato) + " ")
				fsuc.write(str(suc) + " ")
		

	fato.close()
	fsuc.close()
	

	# wva = web3.eth.get_balance("0x9e0c2d4866A5FfF59971749b5E91b5E8fDb8E7cf")
	# print(wva)

def calato1(totalnonce):
	web3 = Web3(Web3.HTTPProvider("http://127.0.0.1:24765"))
	web31 = Web3(Web3.HTTPProvider("http://127.0.0.1:24768"))
	print(web3.isConnected())
	print(web31.isConnected())
	print("....1")


	sender = bytes.fromhex("94abee207bc81159831e3c6b3a6722d019e8e13e")
	sender1 = bytes.fromhex("6f55fcf81624e908da2ba3c83341ff707b652760")

	poster = bytes.fromhex("d3d091e01502d53e3820aed1e6eff3af2673f346")
	poster1 = bytes.fromhex("2d05a3fbe28aadd143701727c8a766694654fb99")

	losingato = 0
	unsuc = 0
	punsuc = 0

	fato = open("fato.txt", "a+")
	fsuc = open("fsuc.txt", "a+")
	fpsuc = open("fpsuc.txt", "a+")

	v = web3.eth.get_transaction_count(to_checksum_address("0xd3d091e01502d53e3820aed1e6eff3af2673f346") )
	v1 = web31.eth.get_transaction_count(to_checksum_address("0x2d05a3fbe28aadd143701727c8a766694654fb99") )



	suc = v1
	unsuc = totalnonce - v1
	losingato = v - v1
	punsuc = totalnonce - v

	fato.write(str(losingato) + " ")
	fsuc.write(str(unsuc) + " ")
	fpsuc.write(str(punsuc) + " ")

	# for i in range(0, totalnonce, 1):
	# 	raw = rlp.encode([sender, i])
	# 	h = keccak(raw)
	# 	address = h[12:]
	# 	# print(to_checksum_address(address) )

	# 	raw1 = rlp.encode([sender1, i])
	# 	h1 = keccak(raw1)
	# 	address1 = h1[12:]
	# 	# print(to_checksum_address(address1) )

	# 	v = web3.eth.get_transaction_count(0xd3d091e01502d53e3820aed1e6eff3af2673f346)
	# 	v1 = web31.eth.get_transaction_count(0x2d05a3fbe28aadd143701727c8a766694654fb99)

	# 	if v > 0 and v1 > 0:
	# 		suc = suc + 1
	# 		ato = ato + 1
	# 		if i == totalnonce -1:
	# 			print(suc, end = ' ')
	# 			print(ato, end = ' ')
	# 			fato.write(str(ato) + " ")
	# 			fsuc.write(str(suc) + " ")
	# 	elif v == 0 and v1 == 0:
	# 		ato = ato + 1
	# 		if i == totalnonce -1:
	# 			print(suc, end = ' ')
	# 			print(ato, end = ' ')
	# 			fato.write(str(ato) + " ")
	# 			fsuc.write(str(suc) + " ")
	# 	else:
	# 		if i == totalnonce -1:
	# 			print(suc, end = ' ')
	# 			print(ato, end = ' ')
	# 			fato.write(str(ato) + " ")
	# 			fsuc.write(str(suc) + " ")
		

	fato.close()
	fsuc.close()
	fpsuc.close()


if __name__ == '__main__':
	parser = argparse.ArgumentParser(description='Process some integers.')

	parser.add_argument('--nonce', '-n', help='an integer for nonce', required = True)

	args = parser.parse_args()



	# print(args.integers[0])

	# nonce = sys.argv[1]


	calato1(int(args.nonce))


