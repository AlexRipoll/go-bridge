import Web3 from "web3";
// import CustodialVault from "../build/contracts/ERC721CustodialVault.json"


const web3 = new Web3(Web3.givenProvider);
// const goerliChainId = "0x5";
// const mumnbaiChainId = "0x13881";
// const bscthainId = "0x61";
//
// function connectMetamask() {
//     ethereum.request({ method: 'eth_requestAccounts' });
// }

const logging = async function() {
    // const connected = await detectEthereumProvider();
    // console.log(">>>> ", connected)
    const networkId = await web3.eth.net.getId()
    console.log(">>>> ", networkId);
}

// async function source() {
//     const accounts = await web3.eth.getAccounts();
//     const account = accounts[0];
//     document.getElementById('wallet-address').textContent = account;
//
//     const connected = await detectEthereumProvider();
//
//     const networkId = await web3.eth.net.getId()
//     console.log(networkId);
//     const networkData = CustodialVault.networks[networkId]
//
//     if (connected.chainId == goerliChainId) {
//         // select contracts and networks
//         var sNft = goeNFT
//         var sCustody = goeCustody
//         var sRpc = goerpc
//         var erc20 = goeErc20
//     }
//     else if (connected.chainId == mm) {
//         var sNft = mumNFT
//         var sCustody = mumCustody
//         var sRpc = mumrpc
//         var erc20 = mumErc20
//     }
//     else if (connected.chainId == bsct) {
//         var sNft = bsctNFT
//         var sCustody = bsctCustody
//         var sRpc = bsctrpc
//         var erc20 = bsctErc20
//     }
//     const provider = new ethers.providers.JsonRpcProvider(sRpc)
//     const key = simpleCrypto.decrypt(cipherEth)
//     const wallet = new ethers.Wallet(key, provider);
//     const contract = new ethers.Contract(sNft, NftABI, wallet);
//     const itemArray = [];
//     await contract.walletOfOwner(account).then((value => {
//         value.forEach(async(id) => {
//             let token = parseInt(id, 16)
//             const rawUri = contract.tokenURI(token)
//             const Uri = Promise.resolve(rawUri)
//             const getUri = Uri.then(value => {
//                 let str = value
//                 let cleanUri = str.replace('ipfs://', 'https://ipfs.io/ipfs/')
//                 let metadata = axios.get(cleanUri).catch(function (error) {
//                     console.log(error.toJSON());
//                 });
//                 return metadata;
//             })
//             getUri.then(value => {
//                 let rawImg = value.data.image
//                 var name = value.data.name
//                 var desc = value.data.description
//                 let image = rawImg.replace('ipfs://', 'https://ipfs.io/ipfs/')
//                 let meta = {
//                     name: name,
//                     img: image,
//                     tokenId: token,
//                     wallet: account,
//                     desc
//                 }
//                 itemArray.push(meta)
//             })
//         })
//     }))
//     await new Promise(r => setTimeout(r, 2000));
//     console.log("Wallet Refreshed : " + sRpc)
//     getSourceNft(sNft);
//     getErc20(erc20);
//     getSourceCustody(sCustody);
//     getSourceRpc(sRpc);
//     setNfts(itemArray);
// }

export default logging();
