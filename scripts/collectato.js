const HELP_MSG =
  'Usage: node TG_worker.js -t [T] -c [N]\n\n\t-c [N]\tDeploy a contract to test.\n\t\tN\tNumber of transactions generated to the contract per second.\n\n\t-t [T]\tSet test duration.\n\t\tT\tTest duration in seconds.\n\n\t -i [I]\tWorker id.\n\t\tI\tId of the worker to be shown in the logs.\n\n\t-h\tDisplay this message.'

const TEST_CONTRACT_1 =
  'contract testContract1{ ' +
  '    struct CallRecord{ ' +
  '        uint callid; ' +
  '        address from; ' +
  '        uint time; ' +
  '    } ' +
  '    CallRecord[] public callRecords; ' +
  '    uint public numberOfCallRecords; ' +
  '    address owner; ' +
  '    constructor() public{ ' +
  '        owner = msg.sender; ' +
  '        numberOfCallRecords = 0; ' +
  '    } ' +
  '    function injectTx(uint callid) public{ ' +
  '        callRecords.push(CallRecord({ ' +
  '            callid: callid, ' +
  '            from: msg.sender, ' +
  '            time: block.timestamp ' +
  '        })); ' +
  '    } ' +
  '} '



const accounts = [
  {
    address: '0xd3d091e01502d53e3820aed1e6eff3af2673f346', //eth1
    privateKey: '0xc17841232ec7c6654ffd79ae3a8bbeec9da63b99fbdf42dffdb820f4d734d0ca'
  },
  {
    address: '0x94abee207bc81159831e3c6b3a6722d019e8e13e', //eth1[1]
    privateKey: '0xa24e4d4657fb5870031d602d0231f14f877ffd8e132af70cd72a938201784466'
  },
  {
    address: '0xb782edeacc48fa6161dc690043bd0db17e959bf4', //ves1[0]
    privateKey: '0x93c4bb70eafb393e6b7d338b0e889de33925935d77fd3fe3540a99c8203378a8'
  },
  {
    address: '0xea2ba9cfae53f04989faefeb6fa7f0b65a72e4b0', //ves1[1]
    privateKey: '0x76e3237d64c57ef50da5d78ff3024942962a439b86ab4cf50bf2c38339ceb1ac'
  },
  {
    address: '0x2d05a3fbe28aadd143701727c8a766694654fb99', //ves2[0]
    privateKey: '0xf2add373ac0979002383c4a41db3bda1278492b4a6764f1a53dc43a73dee4b69'
  },
  {
    address: '0x6f55fcf81624e908da2ba3c83341ff707b652760', //ves2[1]
    privateKey: '0xe45d0d1fc072ed0272cda33c9c2021beb4e4db47e01d471227db6b4ee54cb627'
  },
  {
    address: '0xc157a9260be13546d83bd089450e2cc25b38ae29', //eth2[0]
    privateKey: '0x5625e3268511e22e021a03a713132b0cfcaf47571a8e31a7328129f81a1f5424'
  },
  {
    address: '0xe790a5c30ccaef61bc123326bd3fc65a7f039ba7', //eth2[1]
    privateKey: '5625e3268511e22e021a03a713132b0cfcaf47571a8e31a7328129f81a1f5424'
  }
]

var Web3 = require('web3')
var solc = require('solc')
var fs = require('fs')
var net = require('net')
var path = require('path')


function anotherworker () {
  workerId = 5

  var log = require('./TG_log')

  var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:24768'))
  log.consoleLog(workerId, 'Connected to another Ethereum node.')
  web3.eth
    .getBlockNumber()
    .then(blkNum => log.consoleLog(workerId, 'Block number: ' + blkNum + '.'))
    .catch(() => log.consoleLog(workerId, 'Ethereum node does not reply.'))


  // var keythereum = require("keythereum");
  // var datadir = "/Users/taoyuechen/Library/Ethereum1";
  // var address= "0x94abee207bc81159831e3c6b3a6722d019e8e13e";
  // const password = "1";

  // var keyObject = keythereum.importFromFile(address, datadir);
  // var privateKey = keythereum.recover(password, keyObject);
  // log.consoleLog(workerId, privateKey.toString('hex'));

  var account1 = web3.eth.accounts.privateKeyToAccount(
    accounts[parseInt(workerId)].privateKey
  )
  web3.eth.accounts.wallet.add(account1);

  var account2 = web3.eth.accounts.privateKeyToAccount(
    accounts[parseInt(4)].privateKey
  )

  log.consoleLog(workerId, 'Ethereum account ' + account1.address + ' used.')
  log.consoleLog(workerId, 'Ethereum account private key ' + account1.privateKey)

  const rlp = require('rlp');
  const keccak = require('keccak');

  var nonce = 0x00; //The nonce must be a hex literal!
  var sender = account2.address; //Requires a hex string as input!

  var input_arr = [ sender, nonce ];
  var rlp_encoded = rlp.encode(input_arr);

  var contract_address_long = keccak('keccak256').update(rlp_encoded).digest('hex');

  var contract_address = contract_address_long.substring(24);
  log.consoleLog(workerId, "sender in chain2 "+sender)
  log.consoleLog(workerId, "contract_address in chain 2 "+contract_address)

  web3.eth.getBalance(contract_address)
  .then(blc => {
    if (blc > 0) {
      log.consoleLog(workerId, 'y.....'+blc)
      web3.eth.sendTransaction({
        from: account1.address,
        to: contract_address,
        gas: 215650,
        data: "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000001"
      })
    }
    else {
      log.consoleLog(workerId, 'n.....'+blc)
      web3.eth.sendTransaction({
        from: account1.address,
        to: contract_address,
        gas: 215650,
        data: "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000002"
      })
    }
  });

  // web3.eth.getTransactionCount(accounts[2].address).then(after => 
  //   if (after >= 2) {
  //     log.consoleLog(workerId, after)
  //   }
  //   //todo 
  //   web3.eth.sendTransaction({
  //     from: account.address,
  //     gas: 1000000,
  //     // value: 60000000000000,
  //    // data:  // invoking functions, further conducted
  //   }, function(error, hash){
  //     }
  // )

  
}


function calcu (Contract_array, Contract_array1, totalnonce) {
  for (j = 0; j < totalnonce; j ++){
    // contract_addr = Contract_array[j]
    // contract_addr1 = Contract_array1[j]
    // x = totalnonce-1
    // nonce = j
    // log.consoleLog(workerId, "j is " + JSON.stringify(nonce))
    // log.consoleLog(workerId, 'contract_address is ' + Contract_array[j])
    // web31.eth.getBalance(Contract_array1[j])
    // .then(blc => {log.consoleLog(workerId, 'contract_address after is ' + JSON.stringify(Contract_array1[j]) )
    //   log.consoleLog(workerId, 'BLC is ' + JSON.stringify(blc) )

    // })
    // log.consoleLog(workerId, "contract1 " + Contract_array[j].addr)
    // contract = Contract_array[j].addr

    web3.eth.getBalance( JSON.parse(JSON.stringify(Contract_array[j].addr)) )
    .then(blc => {
      
      // var Contract_array = JSON.parse(Contract_array)
      // log.consoleLog(workerId, 'contract_addresseseeeeee are ' + Contract_array)
      log.consoleLog(workerId, "here")
      // log.consoleLog(workerId, "contract1 " + (Contract_array[j]).addr )

      if(blc > 0) {
        
        // log.consoleLog(workerId, "contract1 " + JSON.stringify(Contract_array[j]) )
        web31.eth.getBalance( JSON.parse(JSON.stringify(Contract_array1[j].addr))  )
        .then(blc1 => {
          log.consoleLog(workerId, "here")
          // log.consoleLog(workerId, 'contract_address after is ' + Contract_array[j])

          if(blc1 > 0) {
            // log.consoleLog(workerId, "suc....."+blc + " "+blc1)
            ato = ato + 1;
            suc = suc + 1;
            log.consoleLog(workerId, "totalnonce is " + JSON.stringify(totalnonce) + ", j is " + JSON.stringify(j) +", atomi is " + JSON.stringify(ato) +", succ is " + JSON.stringify(suc));

            // if(nonce < x-1) {
            //   log.consoleLog(workerId, "totalnonce is " + JSON.stringify(x) + ", i is " + JSON.stringify(nonce) +", atomi is " + JSON.stringify(ato) +", succ is " + JSON.stringify(suc));
            // }
            // else {
            //   log.consoleLog(workerId, "i is " + nonce +", atomi is " + ato +", succ is " + suc);
            //   log.consoleLog(workerId, "contract_address1 is " + contract_address1 +", atomi is " + ato +", succ is " + suc);
            // }

            // if(i == x) {
            //   log.consoleLog(workerId, "i is " + nonce +", atomi is " + ato +", succ is " + suc);
            //   log.consoleLog(workerId, "contract_address1 is " + contract_address1 +", atomi is " + ato +", succ is " + suc);
            // }
          }
          else {
            // log.consoleLog(workerId, "n....."+blc + " "+blc1)
            // log.consoleLog(workerId, "i is " + nonce +" atomi is " + ato +", succ is " + suc);
            if(i == totalnonce - 1) {
              log.consoleLog(workerId, "atomi is " + ato);
              log.consoleLog(workerId, "succe is " + suc);
            }
          }
        })
      }
      else {
        web31.eth.getBalance(Contract_array1[j])
        .then(blc1 => {
          if(blc1 == 0) {
            // log.consoleLog(workerId, "suc" + blc + " " + blc1)
            ato = ato + 1;
            // log.consoleLog(workerId, "i is " + nonce +" atomi is " + ato +", succ is " + suc);
            // if(i == totalnonce - 1) {
            //   log.consoleLog(workerId, "atomi is " + ato);
            //   log.consoleLog(workerId, "succe is " + suc);
            // }
          }
          else {
            // log.consoleLog(workerId, "i is " +i +" atomi is " + ato +", succ is " + suc);
            // if(i == totalnonce - 1) {
            //   log.consoleLog(workerId, "atomi is " + ato);
            //   log.consoleLog(workerId, "succe is " + suc);
            // }
            // log.consoleLog(workerId, "n....."+blc + " "+blc1)
          }
        })
      }
    }

    );
}



function worker (totalnonce) {

  // if (args['c'] !== undefined && !Array.isArray(args['c'])) {
  //   args['c'] = [args['c']]
  // }

  // if (
  //   args['help'] ||
  //   args['h'] ||
  //   args['t'] === undefined ||
  //   isNaN(args['t']) ||
  //   args['c'] === undefined ||
  //   !args['c'].every(i => !isNaN(i))
  // ) {
  //   console.log(HELP_MSG)
  //   return
  // }




  workerId = 0

  var log = require('./TG_log')

  var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:24765'))
  log.consoleLog(workerId, 'Connected to Ethereum node.')
  // web3.eth
  //   .getBlockNumber()
  //   .then(blkNum => log.consoleLog(workerId, 'Block number: ' + blkNum + '.'))
  //   .catch(() => log.consoleLog(workerId, 'Ethereum node does not reply.'))

  var account = web3.eth.accounts.privateKeyToAccount(
    accounts[parseInt(workerId)].privateKey
  )
  web3.eth.accounts.wallet.add(account)

  var account1 = web3.eth.accounts.privateKeyToAccount(
    accounts[parseInt(1)].privateKey
  )

  log.consoleLog(workerId, 'Ethereum account ' + account.address + ' used.')
  log.consoleLog(workerId, 'Ethereum account private key ' + account.privateKey)









  var web31 = new Web3(new Web3.providers.HttpProvider('http://localhost:24768'))
  log.consoleLog(workerId, 'Connected to another Ethereum node.')
  // web31.eth
  //   .getBlockNumber()
  //   .then(blkNum => log.consoleLog(workerId, 'Block number: ' + blkNum + '.'))
  //   .catch(() => log.consoleLog(workerId, 'Ethereum node does not reply.'))

  var account2 = web31.eth.accounts.privateKeyToAccount(
    accounts[parseInt(4)].privateKey
  )
  web31.eth.accounts.wallet.add(account2)

  var account3 = web31.eth.accounts.privateKeyToAccount(
    accounts[parseInt(5)].privateKey
  )

  log.consoleLog(workerId, 'another Ethereum account ' + account2.address + ' used.')
  log.consoleLog(workerId, 'another Ethereum account private key ' + account2.privateKey)


  const rlp = require('rlp');
  const keccak = require('keccak');
  var sender = account1.address;
  var sender1 = account3.address;

  var ato = 0;
  var suc =0;

  var Contract_array = []
  var Contract_array1 = []

  for (i = 0; i < 1; i++) {
    nonce = 10;
    x = totalnonce
    log.consoleLog(workerId, "totalnonce is " + JSON.stringify(x) + ", i is " + JSON.stringify(nonce))

    var input_arr = [ sender, nonce ];
    var rlp_encoded = rlp.encode(input_arr);

    var contract_address_long = keccak('keccak256').update(rlp_encoded).digest('hex');

    var contract_address = contract_address_long.substring(24);



    var input_arr1 = [ sender1, nonce ];
    var rlp_encoded1 = rlp.encode(input_arr1);

    var contract_address_long1 = keccak('keccak256').update(rlp_encoded1).digest('hex');

    var contract_address1 = contract_address_long1.substring(24);

    contract_address = web3.utils.toChecksumAddress(contract_address)
    contract_address1 = web31.utils.toChecksumAddress(contract_address1)

    // Contract_array
    // Contract_array += JSON.stringify(contract_address)
    // Contract_array1 += JSON.stringify(contract_address1)
    var con = '{"addr":' + JSON.stringify(contract_address) + '}'; 
    var con1 = '{"addr":' + JSON.stringify(contract_address1) + '}';
    // con1.addr = contract_address

    Contract_array.push(JSON.parse(con))
    Contract_array1.push(JSON.parse(con1))
    // log.consoleLog(workerId, 'contract_address is ' + JSON.stringify(contract_address))
    // log.consoleLog(workerId, 'contract_address1 is ' + JSON.stringify(contract_address1))


  }

  log.consoleLog(workerId, 'contract_addresses are ' + Contract_array[0].addr)
  log.consoleLog(workerId, 'contract_addresses1 are ' + Contract_array1[0].addr)  

  // return Contract_array, Contract_array1

    // log.consoleLog(workerId, "i is " + nonce +", atomi is " + ato +", succ is " + suc);
}



  // var nonce = 0x00; //The nonce must be a hex literal!
   //Requires a hex string as input!

  


  //var nonce = 0x00; //The nonce must be a hex literal!
   //Requires a hex string as input!

  



  



//initial
  // var ethContractDeployment = ethContract
  //   .deploy({ data: '0x' + bytecode, arguments: [60000000000000] })
  //   .send({ from: account.address, gas: 1000000 })
  //   .then(contract => {
  //     log.consoleLog(
  //       workerId,
  //       'Deployed contract address: ' + contract.options.address
  //     )
  //     web3.eth.getTransactionCount(account.address).then(nonce => {
  //       log.consoleLog(workerId, nonce)
  //       contract.methods
  //       .givetoothers('0x94abee207bc81159831e3c6b3a6722d019e8e13e', 0)
  //       .send({ from: account.address, gas: 1000000, 
  //       nonce: nonce, gasprice: (0 + 1)*200000 })
  //       .on('receipt', recp =>
  //       log.consoleLog(workerId, 'Transaction ' + 0 + ' injected.')
  //       )
  //       // transactionInjector(
  //       //   web3,
  //       //   contract,
  //       //   account,
  //       //   log,
  //       //   workerId,
  //       //   testDuration,
  //       //   contractLoad,
  //       //   nonce,
  //       //   0
  //       // )
  //     })
  //   })
  // .catch(() =>
  //     log.consoleLog(workerId, "Contract deployment failed."));
}


//todo: hyperservice: web3, change to launching smart contracts..., write a script for gettointial, and send transactions to hyperservice; write a contract on hyperservice
function main (argv) {
  args = require('minimist')(argv.slice(2))
  if (args['c'] !== undefined && !Array.isArray(args['c'])) {
    args['c'] = [args['c']]
  }

  if (
    args['help'] ||
    args['h'] ||
    args['t'] === undefined ||
    isNaN(args['t']) ||
    args['c'] === undefined ||
    !args['c'].every(i => !isNaN(i))
  ) {
    console.log(HELP_MSG)
    return
  }

  // var log = require('./TG_log')
  // const rlp = require('rlp');
  // const keccak = require('keccak');

  var nonce = args['t']

  // account1 = accounts[parseInt(0)].address

  // var sender = account1; //Requires a hex string as input!
  // log.consoleLog(0, 'nonce is ' + nonce)

  // var input_arr = [ sender, nonce ];
  // var rlp_encoded = rlp.encode(input_arr);

  // var contract_address_long = keccak('keccak256').update(rlp_encoded).digest('hex');

  // var contract_address = contract_address_long.substring(24);

  // log.consoleLog(0, 'contract address is ' + contract_address)
  worker(nonce)


}

main(process.argv)
// worker(process.argv)
