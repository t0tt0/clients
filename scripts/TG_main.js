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
  var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'))
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
              var fs = require("fs")
              if (args["o"] !== undefined){
                fs.appendFile(args["o"], "Block Number,Timestamp,Number of Transactions\n", (err)=>{})
              }
              console.log("Block Number,Timestamp,Number of Transactions")
              for (var i = startBlk; i <= endBlk; ++i) {
                web3.eth.getBlock(i).then(blk => {
                  var blkStat = {}
                  blkStat.number = blk.number
                  blkStat.timestamp = blk.timestamp
                  blkStat.txCount = blk.transactions.length
                  if (args["o"] !== undefined){
                    fs.appendFile(args["o"], blk.number + ',' + blk.timestamp + ',' + blk.transactions.length + '\n', (err)=>{})
                  }
                  console.log(blk.number + ',' + blk.timestamp + ',' + blk.transactions.length)
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
