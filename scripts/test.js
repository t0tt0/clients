const HELP_MSG =
  'Usage: node TG_main.js -t [T] -c [N] -c [N] -c [N] ...\n\n\t-c [N]\tDeploy a contract to test.\n\t\tN\tNumber of transactions generated to the contract per second.\n\n\t-t [T]\tSet test duration.\n\t\tT\tTest duration in seconds. (Default: 180)\n\n\t-h\tDisplay this message.'

const workerId = 999


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

  var test_time = args['t']
  var contract_load = args['c']

  if (
    isNaN(test_time) ||
    (Array.isArray(contract_load) && !contract_load.every(i => !isNaN(i))) ||
    (!Array.isArray(contract_load) && isNaN(contract_load))
  ) {
  }

  var child_process = require('child_process')
  var path = require('path')

  var Web3 = require('web3')
  var log = require('./TG_log')
  var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:24765'))
  log.consoleLog(workerId, 'Connected to Ethereum node.')
  var startBlk = 0
  var endBlk = 99999
  web3.eth
    .getBlockNumber()
    .then(blkNum => {
      log.consoleLog(workerId, 'Block number: ' + blkNum + '.')
      startBlk = blkNum
    })
    .catch(() => log.consoleLog(workerId, 'Ethereum node does not reply.'))

  var counter = 0

  contract_load.forEach(i => {
    child_process
      .fork(path.join(path.dirname(argv[1]), 'TG_worker.js'), [
        '-t',
        test_time,
        '-c',
        i,
        '-i',
        (counter++).toString()
      ])
      .on('exit', () => {
        counter--
        if (counter === 0) {
          web3.eth
            .getBlockNumber()
            .then(blkNum => {
              log.consoleLog(workerId, 'Block number: ' + blkNum + '.')
              endBlk = blkNum
              for (var i = startBlk; i <= endBlk; ++i) {
                web3.eth.getBlock(i).then(blk => {
                  var blkStat = {}
                  blkStat.number = blk.number
                  blkStat.timestamp = blk.timestamp
                  blkStat.txCount = blk.transactions.length
                  log.consoleLog(
                    workerId,
                    'Block: ' + JSON.stringify(blkStat) + '.'
                  )
                })
              }
            })
            .catch(() =>
              log.consoleLog(workerId, 'Ethereum node does not reply.')
            )
        }
      })
  })
}

main(process.argv)

// var Web3 = require('web3')
// var Personal = require('web3-eth-personal')
// var log = require('./TG_log')
//   // var web3 = new Web3(Web3.givenProvider || 'http://localhost:24765');
// var personal = new Personal(new Personal.providers.HttpProvider('http://localhost:24765'))
// var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:24765'))


//  function gettime() {

//   // web3.eth.getBlock(2000).then(msg => {console.log(msg)})

//     var keyth=require('keythereum');
//     var keyobj=keyth.importFromFile('0x94abee207bc81159831e3c6b3a6722d019e8e13e','/Users/taoyuechen/Library/Ethereum1/');
//     var privateKey=keyth.recover('1', keyobj);
//     console.log(privateKey.toString('hex'));




//     web3.eth
//     .getBlockNumber()
//     .then(blkNum => {
//       log.consoleLog(workerId, 'the lastest Block number: ' + blkNum + '.');
//       blk = blkNum
//       console.log(blk)
//       for (blk = blkNum; blk > 568; blk =blk - 1) {
//         web3.eth.getBlock(blk).then(tmp => {
//           if(tmp['transactions'] != false) {
//             x = (tmp['timestamp'] - 1619091526.5980318)
//             log.consoleLog(workerId, 'timestamp: ' + x + '    blockid    '+ tmp['number'])
//             return;
//           }
//         })
//       }
//     })
//     .catch(() => log.consoleLog(workerId, 'Ethereum node.'))

//     web3.eth.getTransactionCount("0xd3d091e01502d53e3820aed1e6eff3af2673f346").then(nonce => {log.consoleLog(workerId, 'number of transactions currently: ' + nonce)
//   })
// }

// //nonce 58

// //4 177.8911862373352


// // 8 160.82454442977905
// // 9 225.0979881286621
// // 10 189.2469916343689
// // 15 288.41501808166504
// // 13 169.46155405044556
// // 10 230.56667184829712
// // 16 346.77179193496704
// // 16 285.0786519050598
// // 16 228.98546028137207

// //128 




// //

// function main (argv) {
  //gas limit 0x5210 before
  
//   log.consoleLog(workerId, 'Connected to Ethereum node.')
//   log.consoleLog(workerId, 'Connected to Ethereum node.')
//   // var startBlk = 0
//   // var endBlk = 99999


//  gettime()

  

  
// }

// main(process.argv)
