const bitcoin = require('bitcoinjs-lib')

var NETWORK = {
  messagePrefix: '\x18Bitcoin Signed Message:\n',
  bech32: 'bc',
  bip32: {
    public: 0x0488b21e,
    private: 0x0488ade4
  },
  pubKeyHash: 0x00,
  scriptHash: 0x05,
  wif: 0x80
}

// var keyPair = bitcoin.ECPair.fromWIF(wif, NETWORK)
var keyPair = bitcoin.ECPair.makeRandom()
if (!keyPair.compressed) {
  throw new Error('Segwit supports only compressed public keys')
}
var pubKey = keyPair.getPublicKeyBuffer()
var pubKeyHash = bitcoin.crypto.hash160(pubKey)

var redeemScript = bitcoin.script.witnessPubKeyHash.output.encode(pubKeyHash)
var redeemScriptHash = bitcoin.crypto.hash160(redeemScript)
var scriptPubKey = bitcoin.script.scriptHash.output.encode(redeemScriptHash)
var segwit = bitcoin.address.fromOutputScript(scriptPubKey, NETWORK)
console.log("common address:", keyPair.getAddress())
console.log('Segwit Address:', segwit)
console.log("private key:", keyPair.toWIF())

