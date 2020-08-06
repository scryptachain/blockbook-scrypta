// +build unittest

package liquid

import (
	"encoding/hex"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/scryptachain/blockbook-scrypta/bchain"
	"github.com/scryptachain/blockbook-scrypta/bchain/coins/btc"
)

func TestMain(m *testing.M) {
	c := m.Run()
	chaincfg.ResetParams()
	os.Exit(c)
}

func Test_GetAddrDescFromAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "P2PKH1",
			args:    args{address: "QHU1yszeZwVeuJosGJ4JDHuKaLRWmdEYDF"},
			want:    "76a914dd95db91e8f914cbd63bae8e307d54399f060cd688ac",
			wantErr: false,
		},
		{
			name:    "P2PKH2",
			args:    args{address: "PwogNb9zqhrwPneDqVp18GqcVuygsCvXkU"},
			want:    "76a91405eb4afe4615751cfb813a00846a8d9ef8a9a2e588ac",
			wantErr: false,
		},
		{
			name:    "P2SH1",
			args:    args{address: "GhWTZqLPHRK8KfuT6yo1wGisQzn4cXrbPP"},
			want:    "a9140394b3cf9a44782c10105b93962daa8dba304d7f87",
			wantErr: false,
		},
	}
	parser := NewLiquidParser(GetChainParams("main"), &btc.Configuration{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.GetAddrDescFromAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddrDescFromAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("GetAddrDescFromAddress() = %v, want %v", h, tt.want)
			}
		})
	}
}

func Test_GetAddressesFromAddrDesc(t *testing.T) {
	type args struct {
		script string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		want2   bool
		wantErr bool
	}{
		{
			name:    "P2PKH1",
			args:    args{script: "76a914dd95db91e8f914cbd63bae8e307d54399f060cd688ac"},
			want:    []string{"QHU1yszeZwVeuJosGJ4JDHuKaLRWmdEYDF"},
			want2:   true,
			wantErr: false,
		},
		{
			name:    "P2PKH2",
			args:    args{script: "76a91405eb4afe4615751cfb813a00846a8d9ef8a9a2e588ac"},
			want:    []string{"PwogNb9zqhrwPneDqVp18GqcVuygsCvXkU"},
			want2:   true,
			wantErr: false,
		},
		{
			name:    "P2SH",
			args:    args{script: "a9140394b3cf9a44782c10105b93962daa8dba304d7f87"},
			want:    []string{"GhWTZqLPHRK8KfuT6yo1wGisQzn4cXrbPP"},
			want2:   true,
			wantErr: false,
		},
		{
			name:    "P2PK compressed",
			args:    args{script: "21020e46e79a2a8d12b9b5d12c7a91adb4e454edfae43c0a0cb805427d2ac7613fd9ac"},
			want:    []string{"QKKEbCNAV7BYdtfSb3cQcJ7QSmFMvtXETz"},
			want2:   false,
			wantErr: false,
		},
		{
			name:    "P2PK uncompressed",
			args:    args{script: "41041057356b91bfd3efeff5fc0fa8b865faafafb67bd653c5da2cd16ce15c7b86db0e622c8e1e135f68918a23601eb49208c1ac72c7b64a4ee99c396cf788da16ccac"},
			want:    []string{"QDoUiWY7iZXDrkBzXdk6dru8DbvGqExXuf"},
			want2:   false,
			wantErr: false,
		},
	}

	parser := NewLiquidParser(GetChainParams("main"), &btc.Configuration{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := hex.DecodeString(tt.args.script)
			got, got2, err := parser.GetAddressesFromAddrDesc(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddressesFromAddrDesc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddressesFromAddrDesc() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetAddressesFromAddrDesc() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

var (
	testTx1 bchain.Tx

	testTxPacked1 = "0a207aa1af9481f2d744c96015b1baea6ba753790971ff265adfd93775de0234bfd612b51a020000000101a99547d213b005f355da348de54f5eb370fbc6a5687e412897ef0ec4ce237d75020000006b483045022100ae926c96c746308e7488e022f4ad1db94d5d0c8683f6fa6ded3afb13d8e20578022074aa8ebfe20adaf25beed70c60cfb5007278bec17875f11f7a5b3c86eb60a96901210391abdcd113c40b56f13a548e8624d6ad8d7a162b33ef81020d046cadbda26637feffffff030af62b535fc393152f6d575708b271e3a53514cdcf65508c7d12e8fd06709ce24208e192e1315aa94c3bc0eddbb0ae195aff0a1ba2e19773b79e5d1b29fdd8df211b02f8b255b83fa13f40745fb5054b89dc7dcba497c850086a8726bac66ac22d4be31976a914d79c1c8b67a0275c60e33b67bbd0e19a79b9276388ac016d521c38ec1ea15734ae22b7c46064412829c0d0579f0a713d1c04ede979026f01000000000000000000256a23426c6f636b73747265616d2e696e666f206973206120636f6f6c206578706c6f726572016d521c38ec1ea15734ae22b7c46064412829c0d0579f0a713d1c04ede979026f0100000000000004800000000000000000000043010001db986007e38ccb7bdef1661fcf633cbde3d8850cd116736d9b0b0b027901921694954b26036eb51d9424c38083aa3d937998ea81077843925513523eb6ec3337fd4d0b602300000000000000011486003b49c8ab1a6aa09a8eaec4ed83dbda4cb4ef9651114513823a6eefb1dbf971631d88a229a3b027a94f9f15966765f6f5ce3245057c57aedf8f4c69e355f4264a10e2c7dad691879844074ae5e37152c828aba89035e928c63b1c1590638333d05ed08442538ebf7dff9c94ce11ab6bb6d5c9535fce99fab023aabc7215b52eed15454333428b065fb5d8b03aa2fe1de5004f9d8fca717bcd682f1f9caa6561bdafbe69c8423166f7e33867f0bc4fdd85224ff1533839762530d47a1f053c40c7e38f84a1431ad03398bc9d634384aec0f22e75ac4a94e3703ec8715d0791564a8b1509eab4bf7543ef2e5c20fe9fb6c2a85afe42fcdba64c628103ea5693287a28ded79f517a3e877fe287ca6f3d1229295abaf5c1b3b43ab54ec4697c941da47e0934718b697a414d8fd1a722eb23ceb554afb1b4c807a94507a35b153f19da10c6688c71efa0d1ad15c8ac2f3d786ae8c43136cacf333c2643e521322314346b201427f1ac975340cf3caf102245d8cf45709bb51e41fd357f1fb7316138f1992e5646842eb4fe18ef6a7096f7e4c1e69950285f30c64dc1da661c055944356f140d6298f5eaf733799bef034a8afb05f6f74239e572acf5d1c2f9e028b806caab13eb19148456b3db0f719ad76818eaea11e7989e74d858743ffba738a5f5c7d12ab1b049c594656821620e20817c6acc5b3d3725f9183d3c01a0190d5d991d1603682418f4a9c55b59cab8787f463123e1efd7d3e7f75dfb9ccf931ae2dc965b82cabf2af6b93293a4b00d7145d97a24679076157bade5ba4d7577922052c719ebf2493f5e9d0cf8e6f82666e8732dd91068395e528d0ea532e27ae8a84ed516dabd30dd990a1539d93f7b775039d53f1a12d7e7747a64f6aca9d7589330c8adf8854dce0fee0be212716d6c1b075fabb9d0c815e613c772f7c57e7efdb4e64a5ef47c9cf384deab55191d7b4b0cb27986649eca2c4286a01a7d3c304f008c6d3fe53abf320afc7602654cf4b69b87fd2b9b6794fe7cc77e92bbef572dc8ddb1530da6d82268cd32db112afb8c2db69651959c1ad39d804178cae05f856539559e3dd0e869d9985f41262c30058d1d3c2cd336649b3a893bd4b29c206c53eeb337025a7585fc4bc05d8fe71035ea082ea543cbc15f8e6f0658815c7553795917958b503779cf6a18f92c391769e846327d3cd8458c089cb1e342a590eb57a583365c4bd5de34dc17992f5309fb69b67aa8436c763183a8f11a543fdb8a376a30836013fff3a2d21b72e22f1d9d031a3cd467365256e120894e1489c238df5ce65183b93d68024697245be8c0a9d60e31452d247a98136ca559622e2f4b35569ec8ab8539528a09214f1d3862bc97b21b7bb2b66c654ebee00eb26ec57988ff174cab65d4864eea14f15b65252fe534750bea3b9b4400c811ccf3e6517ecf1f1585d908f21b91bfd223eddbb9b980bb65133934c8027fde78865ed4d969bc6de3be613730d3a0f179f0ed734c67900da53c230045d60d31980401e7fa1396891d42555af8152bb0e6557d3c7d718f7ba85b6848ab052fcf709c7aea676d8c00c787a7fb2a2b1972d245aa51a24ceb13d7dd38cd0c214727a1000561828a62c969afa60d36cfab21b8e6837a0e0ac325c6c20e3ccd09fd1f5400dd6a7dc861e050bcb47ee89670d807b514e5b59ea830ef885ac6a0371efdcbb09e676fd9790e00ab9686f59b2ce173896d4d451e6fd61f3b95b82e63a494d7853c78c2f03e45cbabd46a4586364c131000d9efdc71d30d3caf25cdfad63a6628e67d5f219a0cafc5509f556cec7ac110ce0b52b552be6780e8d0c31a068dc6607ebc9cac4471cb1e85e5b6a0bc2330062f6930e2ef623da059da02902fa28586614a053625ea5901dc6151e7fae3a63b444b9d53630dea6b90d3c2ca7dd8db69f39a2ad6bc2eec08a85c1c0d0a9b079825278cf0b6f415344b0b6e4c7a51947e98ada9149c8f427914d8245b152f8f3558178d16e35498649a34ceb2fdaafc0d303829ddd9412c9a2b5ed1ce060472a9a85ffa19a4ddfbdb89437e72b261472d2b8c70a01a96562b753692c7329d75057a9918dbb3a39a90c86b66ec745a14e9909b2eae6c46dcd8c8a666973ba356124d6444353f1a34f88f907af897ea8e642f8603aed400d3d8a75568baa96b26b8d04f5fb0d1ca1256e7057f93464d95d1994ac5189ca607c013637fa35879a6c67c4806b6c9f00ccedb953103bbec21f054d304b621e0eca771accf5181409642d58ea8d032e06c27295e57092b4e3a89f47f47f16dc82f07dfad44ca7077385f62b4ad2efecb97955759c31977719316545b6fc6a78b7e35206719fcba4c1dca5d0f7f9959f7bd1c117532a2f7fc3ca87820e38cfab558dc48adb1058964b6ea9dfde5e06bff6b5d7af19fe6b8a46d0ec76f8902e525082591b82c2f86a30341787a1da86ab517063b39ae076b3f10e78f6c79dd58041b99ab40824b598f6c815733d62d262abeb96b8224b56e4109d4d5c445cd055a9cd2aea7753fc0f6a55bcf4291551a0d9c852e19e0f0603eefc9a79a38bc07a73f6e0f88a9d3eac01ca57d5d55e2ef870743522c15841b4ea7a0d278922cc977724571c75b5039a20f5dad843941d91fb421944b54923ed8c5f71edafbdfaf4c12dbe1c2b2a9642bbc2bd07017a44cb402d6ae09248eb24b70d43f751d3bf761ecef51f1f15239de25222852f95607335b3410050e5bd65f8c9657f0a95185e48dcb2f711641c0994e352b0039d569ddbd27182c263649d3e27b3816acd5822b31c98661adf01ee62db361be1b0469de4fc1c6e10185203eb44abe3d077fcbc52bef6095ca30c81a22b6e7a64c1385e7dd11209fc073a29ef656769fbb24d4292d98379884a1a2b2c7f6fd895577d739a2f0411e8afcee4de40927443f28e72aebc763fd0315ee85d29a23bb235ce82a5a621c0d9021965673720be9057a26b1ff0380263b54777f6cc7b7531d284f5041d4480108e4b4986f1ddec9e65dd2392c7fa0349c2a39954e4aeec8fe8f40e66db57d29beaf0046d9387482bab71a9397d90611cf767637c8666da87d5f1d798558cbb7228844010510cb95b073cec3878893ba70549eb6d3428b8db6944118f6de2ee7107b593ef85441cba46238f7843c4e6e7497f14cc64c25653c87da756226ce774c2b5e43294f2ebab2f601e9fe3f2d1d3172cfcc7e6eb7ea9b237e5f093443b02f42b4ea85673a6a0000ef9a6ddf2263a1a75eb7b78d3f90b1de91a22d78aa06d9626b0f9090ee63d92418b084db647132e3b0b7f3ed583eec280d06b94a358daffcd333233fb390ba8bf2da3921b18adf6cd4901cbabf4f4e3c90f21eb3190c8c0e4f16ac25dd546859bc640354ab9769553aeda466191ba4b10f52da5342347685e52af5d20ba8c113e65663bcedc12c99e576c4e1bdde013017d16fc26e3f30418c21d72ad6507ef1c8d437342d0fc20ad102e6c49eb9a8e7a3df5366c9a75b6d95ab007a3d93bca0086414ac5bc44a872659f43f0f703b415ac0e9aeeedb2c0cb945938923dc0865c5d3ff673e068d2865b11c68774cd6c0be1caa40627425bdc4ecbbd0a642f9c6953464e2f20681994be8483d64ed6d2d8efa79c5b12776900e58b45bea18c2e0220ea27e9670485e2c6ce9f52ac08cad61ca57b839710209db8a5d4fe95a846960b126271e519545e5d4d15300fc0ef2f35b8def7ba6f639d85404b57bac35140bef1c4f3b455773bf2a2eae62118dcdc5474ec0900aca49300833417786d1fb0d76a56570cfa10d23dbab0305e1c9c48032b19fd2ec2e00b1528a248036590e26e8c75209d004e20bc7730b29cf3ed860848b83ab91ad6a635f2cc1eca89e16814f34f2c1c2766a28b2901170bb4839f08f5685e593b8a5ce2a801194818b4aeb0a01794f92c7bc4144808e997cfd1711485b60483cf4a310b23e0210c5c73e6956b9e1ae696f1ea79b3fab788bf229a349fd9caedebd5db99821b1cebdceafb011cfcd1c78f93774b35f25c7e69952b0bca1a0d40791b812614e996ca31548bc0000000018cfdaa0e0052000288781053299010a001220757d23cec40eef9728417e68a5c6fb70b35e4fe58d34da55f305b013d24795a91802226b483045022100ae926c96c746308e7488e022f4ad1db94d5d0c8683f6fa6ded3afb13d8e20578022074aa8ebfe20adaf25beed70c60cfb5007278bec17875f11f7a5b3c86eb60a96901210391abdcd113c40b56f13a548e8624d6ad8d7a162b33ef81020d046cadbda2663728feffffff0f3a4110001a1976a914d79c1c8b67a0275c60e33b67bbd0e19a79b9276388ac222251477652524658424d32666558755533394577484e75595953726e4d33486155794d3a2910011a256a23426c6f636b73747265616d2e696e666f206973206120636f6f6c206578706c6f7265723a060a02048010024002"
)

func init() {
	testTx1 = bchain.Tx{
		Hex:       "020000000101a99547d213b005f355da348de54f5eb370fbc6a5687e412897ef0ec4ce237d75020000006b483045022100ae926c96c746308e7488e022f4ad1db94d5d0c8683f6fa6ded3afb13d8e20578022074aa8ebfe20adaf25beed70c60cfb5007278bec17875f11f7a5b3c86eb60a96901210391abdcd113c40b56f13a548e8624d6ad8d7a162b33ef81020d046cadbda26637feffffff030af62b535fc393152f6d575708b271e3a53514cdcf65508c7d12e8fd06709ce24208e192e1315aa94c3bc0eddbb0ae195aff0a1ba2e19773b79e5d1b29fdd8df211b02f8b255b83fa13f40745fb5054b89dc7dcba497c850086a8726bac66ac22d4be31976a914d79c1c8b67a0275c60e33b67bbd0e19a79b9276388ac016d521c38ec1ea15734ae22b7c46064412829c0d0579f0a713d1c04ede979026f01000000000000000000256a23426c6f636b73747265616d2e696e666f206973206120636f6f6c206578706c6f726572016d521c38ec1ea15734ae22b7c46064412829c0d0579f0a713d1c04ede979026f0100000000000004800000000000000000000043010001db986007e38ccb7bdef1661fcf633cbde3d8850cd116736d9b0b0b027901921694954b26036eb51d9424c38083aa3d937998ea81077843925513523eb6ec3337fd4d0b602300000000000000011486003b49c8ab1a6aa09a8eaec4ed83dbda4cb4ef9651114513823a6eefb1dbf971631d88a229a3b027a94f9f15966765f6f5ce3245057c57aedf8f4c69e355f4264a10e2c7dad691879844074ae5e37152c828aba89035e928c63b1c1590638333d05ed08442538ebf7dff9c94ce11ab6bb6d5c9535fce99fab023aabc7215b52eed15454333428b065fb5d8b03aa2fe1de5004f9d8fca717bcd682f1f9caa6561bdafbe69c8423166f7e33867f0bc4fdd85224ff1533839762530d47a1f053c40c7e38f84a1431ad03398bc9d634384aec0f22e75ac4a94e3703ec8715d0791564a8b1509eab4bf7543ef2e5c20fe9fb6c2a85afe42fcdba64c628103ea5693287a28ded79f517a3e877fe287ca6f3d1229295abaf5c1b3b43ab54ec4697c941da47e0934718b697a414d8fd1a722eb23ceb554afb1b4c807a94507a35b153f19da10c6688c71efa0d1ad15c8ac2f3d786ae8c43136cacf333c2643e521322314346b201427f1ac975340cf3caf102245d8cf45709bb51e41fd357f1fb7316138f1992e5646842eb4fe18ef6a7096f7e4c1e69950285f30c64dc1da661c055944356f140d6298f5eaf733799bef034a8afb05f6f74239e572acf5d1c2f9e028b806caab13eb19148456b3db0f719ad76818eaea11e7989e74d858743ffba738a5f5c7d12ab1b049c594656821620e20817c6acc5b3d3725f9183d3c01a0190d5d991d1603682418f4a9c55b59cab8787f463123e1efd7d3e7f75dfb9ccf931ae2dc965b82cabf2af6b93293a4b00d7145d97a24679076157bade5ba4d7577922052c719ebf2493f5e9d0cf8e6f82666e8732dd91068395e528d0ea532e27ae8a84ed516dabd30dd990a1539d93f7b775039d53f1a12d7e7747a64f6aca9d7589330c8adf8854dce0fee0be212716d6c1b075fabb9d0c815e613c772f7c57e7efdb4e64a5ef47c9cf384deab55191d7b4b0cb27986649eca2c4286a01a7d3c304f008c6d3fe53abf320afc7602654cf4b69b87fd2b9b6794fe7cc77e92bbef572dc8ddb1530da6d82268cd32db112afb8c2db69651959c1ad39d804178cae05f856539559e3dd0e869d9985f41262c30058d1d3c2cd336649b3a893bd4b29c206c53eeb337025a7585fc4bc05d8fe71035ea082ea543cbc15f8e6f0658815c7553795917958b503779cf6a18f92c391769e846327d3cd8458c089cb1e342a590eb57a583365c4bd5de34dc17992f5309fb69b67aa8436c763183a8f11a543fdb8a376a30836013fff3a2d21b72e22f1d9d031a3cd467365256e120894e1489c238df5ce65183b93d68024697245be8c0a9d60e31452d247a98136ca559622e2f4b35569ec8ab8539528a09214f1d3862bc97b21b7bb2b66c654ebee00eb26ec57988ff174cab65d4864eea14f15b65252fe534750bea3b9b4400c811ccf3e6517ecf1f1585d908f21b91bfd223eddbb9b980bb65133934c8027fde78865ed4d969bc6de3be613730d3a0f179f0ed734c67900da53c230045d60d31980401e7fa1396891d42555af8152bb0e6557d3c7d718f7ba85b6848ab052fcf709c7aea676d8c00c787a7fb2a2b1972d245aa51a24ceb13d7dd38cd0c214727a1000561828a62c969afa60d36cfab21b8e6837a0e0ac325c6c20e3ccd09fd1f5400dd6a7dc861e050bcb47ee89670d807b514e5b59ea830ef885ac6a0371efdcbb09e676fd9790e00ab9686f59b2ce173896d4d451e6fd61f3b95b82e63a494d7853c78c2f03e45cbabd46a4586364c131000d9efdc71d30d3caf25cdfad63a6628e67d5f219a0cafc5509f556cec7ac110ce0b52b552be6780e8d0c31a068dc6607ebc9cac4471cb1e85e5b6a0bc2330062f6930e2ef623da059da02902fa28586614a053625ea5901dc6151e7fae3a63b444b9d53630dea6b90d3c2ca7dd8db69f39a2ad6bc2eec08a85c1c0d0a9b079825278cf0b6f415344b0b6e4c7a51947e98ada9149c8f427914d8245b152f8f3558178d16e35498649a34ceb2fdaafc0d303829ddd9412c9a2b5ed1ce060472a9a85ffa19a4ddfbdb89437e72b261472d2b8c70a01a96562b753692c7329d75057a9918dbb3a39a90c86b66ec745a14e9909b2eae6c46dcd8c8a666973ba356124d6444353f1a34f88f907af897ea8e642f8603aed400d3d8a75568baa96b26b8d04f5fb0d1ca1256e7057f93464d95d1994ac5189ca607c013637fa35879a6c67c4806b6c9f00ccedb953103bbec21f054d304b621e0eca771accf5181409642d58ea8d032e06c27295e57092b4e3a89f47f47f16dc82f07dfad44ca7077385f62b4ad2efecb97955759c31977719316545b6fc6a78b7e35206719fcba4c1dca5d0f7f9959f7bd1c117532a2f7fc3ca87820e38cfab558dc48adb1058964b6ea9dfde5e06bff6b5d7af19fe6b8a46d0ec76f8902e525082591b82c2f86a30341787a1da86ab517063b39ae076b3f10e78f6c79dd58041b99ab40824b598f6c815733d62d262abeb96b8224b56e4109d4d5c445cd055a9cd2aea7753fc0f6a55bcf4291551a0d9c852e19e0f0603eefc9a79a38bc07a73f6e0f88a9d3eac01ca57d5d55e2ef870743522c15841b4ea7a0d278922cc977724571c75b5039a20f5dad843941d91fb421944b54923ed8c5f71edafbdfaf4c12dbe1c2b2a9642bbc2bd07017a44cb402d6ae09248eb24b70d43f751d3bf761ecef51f1f15239de25222852f95607335b3410050e5bd65f8c9657f0a95185e48dcb2f711641c0994e352b0039d569ddbd27182c263649d3e27b3816acd5822b31c98661adf01ee62db361be1b0469de4fc1c6e10185203eb44abe3d077fcbc52bef6095ca30c81a22b6e7a64c1385e7dd11209fc073a29ef656769fbb24d4292d98379884a1a2b2c7f6fd895577d739a2f0411e8afcee4de40927443f28e72aebc763fd0315ee85d29a23bb235ce82a5a621c0d9021965673720be9057a26b1ff0380263b54777f6cc7b7531d284f5041d4480108e4b4986f1ddec9e65dd2392c7fa0349c2a39954e4aeec8fe8f40e66db57d29beaf0046d9387482bab71a9397d90611cf767637c8666da87d5f1d798558cbb7228844010510cb95b073cec3878893ba70549eb6d3428b8db6944118f6de2ee7107b593ef85441cba46238f7843c4e6e7497f14cc64c25653c87da756226ce774c2b5e43294f2ebab2f601e9fe3f2d1d3172cfcc7e6eb7ea9b237e5f093443b02f42b4ea85673a6a0000ef9a6ddf2263a1a75eb7b78d3f90b1de91a22d78aa06d9626b0f9090ee63d92418b084db647132e3b0b7f3ed583eec280d06b94a358daffcd333233fb390ba8bf2da3921b18adf6cd4901cbabf4f4e3c90f21eb3190c8c0e4f16ac25dd546859bc640354ab9769553aeda466191ba4b10f52da5342347685e52af5d20ba8c113e65663bcedc12c99e576c4e1bdde013017d16fc26e3f30418c21d72ad6507ef1c8d437342d0fc20ad102e6c49eb9a8e7a3df5366c9a75b6d95ab007a3d93bca0086414ac5bc44a872659f43f0f703b415ac0e9aeeedb2c0cb945938923dc0865c5d3ff673e068d2865b11c68774cd6c0be1caa40627425bdc4ecbbd0a642f9c6953464e2f20681994be8483d64ed6d2d8efa79c5b12776900e58b45bea18c2e0220ea27e9670485e2c6ce9f52ac08cad61ca57b839710209db8a5d4fe95a846960b126271e519545e5d4d15300fc0ef2f35b8def7ba6f639d85404b57bac35140bef1c4f3b455773bf2a2eae62118dcdc5474ec0900aca49300833417786d1fb0d76a56570cfa10d23dbab0305e1c9c48032b19fd2ec2e00b1528a248036590e26e8c75209d004e20bc7730b29cf3ed860848b83ab91ad6a635f2cc1eca89e16814f34f2c1c2766a28b2901170bb4839f08f5685e593b8a5ce2a801194818b4aeb0a01794f92c7bc4144808e997cfd1711485b60483cf4a310b23e0210c5c73e6956b9e1ae696f1ea79b3fab788bf229a349fd9caedebd5db99821b1cebdceafb011cfcd1c78f93774b35f25c7e69952b0bca1a0d40791b812614e996ca31548bc00000000",
		Blocktime: 1544039759,
		Time:      1544039759,
		Txid:      "7aa1af9481f2d744c96015b1baea6ba753790971ff265adfd93775de0234bfd6",
		LockTime:  0,
		Version:   2,
		Vin: []bchain.Vin{
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "483045022100ae926c96c746308e7488e022f4ad1db94d5d0c8683f6fa6ded3afb13d8e20578022074aa8ebfe20adaf25beed70c60cfb5007278bec17875f11f7a5b3c86eb60a96901210391abdcd113c40b56f13a548e8624d6ad8d7a162b33ef81020d046cadbda26637",
				},
				Txid:     "757d23cec40eef9728417e68a5c6fb70b35e4fe58d34da55f305b013d24795a9",
				Vout:     2,
				Sequence: 4294967294,
			},
		},
		Vout: []bchain.Vout{
			{
				N: 0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914d79c1c8b67a0275c60e33b67bbd0e19a79b9276388ac",
					Addresses: []string{
						"QGvRRFXBM2feXuU39EwHNuYYSrnM3HaUyM",
					},
				},
			},
			{
				ValueSat: *big.NewInt(0),
				N:        1,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "6a23426c6f636b73747265616d2e696e666f206973206120636f6f6c206578706c6f726572",
				},
			},
			{
				ValueSat: *big.NewInt(1152),
				N:        2,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "",
				},
			},
		},
	}
}

func Test_PackTx(t *testing.T) {
	type args struct {
		tx        bchain.Tx
		height    uint32
		blockTime int64
		parser    *LiquidParser
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "liquid-1",
			args: args{
				tx:        testTx1,
				height:    82055,
				blockTime: 1544039759,
				parser:    NewLiquidParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    testTxPacked1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.parser.PackTx(&tt.args.tx, tt.args.height, tt.args.blockTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("packTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("packTx() = %v, want %v", h, tt.want)
			}
		})
	}
}

func Test_UnpackTx(t *testing.T) {
	type args struct {
		packedTx string
		parser   *LiquidParser
	}
	tests := []struct {
		name    string
		args    args
		want    *bchain.Tx
		want1   uint32
		wantErr bool
	}{
		{
			name: "liquid-1",
			args: args{
				packedTx: testTxPacked1,
				parser:   NewLiquidParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    &testTx1,
			want1:   82055,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := hex.DecodeString(tt.args.packedTx)
			got, got1, err := tt.args.parser.UnpackTx(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpackTx() got = %+v, want %+v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("unpackTx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
