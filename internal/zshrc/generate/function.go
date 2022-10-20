package generate

const (
	UtilFunctions = `
function dme() {
	mkdir -p ${HOME}/Desktop/deleteme
	cd ${HOME}/Desktop/deleteme
}
`
	ChromeFunctions = `
function chrome() {
    if [[ "${1}" != "" ]]; then
      /usr/bin/open -a "/Applications/Google Chrome.app" ${1}
    fi
}
`
	BlockChainFunctions = `
function btca() {
    chrome "https://www.blockchain.com/btc/address/${1}"
}

function btct() {
    chrome "https://www.blockchain.com/btc/tx/${1}"
}

function btctc() {
    chrome "http://api.blockcypher.com/v1/btc/main/txs/${1}"
}
`

	GodaddyFunctions = `
function domaindaddy() {
    chrome "https://dcc.godaddy.com/control/${1}/dns"
}
`

	LoadTestFunctions = `
function bombard(){
  sleep_duration=${BOMBARD_SLEEP:-2}
  while true;
  do
    date
    curl ${*}
    echo "end-of-response"
    sleep ${sleep_duration}
  done
}


function bombard_grpc(){
  sleep_duration=${BOMBARD_SLEEP:-2}
  while true;do ./grpcurl -plaintext master.lbn-java-grpc-client:80 HostCli/GetHostInfo | jq .hostname;done
}
`
	LocalHostFunction = `
function lh() {
  chrome "http://localhost:${1}"
}
`
)
