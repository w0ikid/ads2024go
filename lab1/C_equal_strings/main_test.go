package main

import "testing"

func Test_compare(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s1: "abc##",
				s2: "a#b#a",
			},
			want: "Yes",
		},
		{
			name: "test2",
			args: args{
				s1: "ab#c",
				s2: "ad#c",
			},
			want: "Yes",
		},
		{
			name: "test3",
			args: args{
				s1: "a#c",
				s2: "bb##",
			},
			want: "No",
		},
		{
			name: "test4",
			args: args{
				s1: "ab##",
				s2: "c#d#",
			},
			want: "Yes",
		},
		{
			name: "test5",
			args: args{
				s1: "rk#ufk#qnt#fzmlg#el#pbldr#rdxuqe#atdhsitbffldp##dkrqomxpffelb##kyqqjt#slet#jkgq##xcpkh##df#y#yqjm##bfuz#fac###yle#wrhxr##hjd#lbwco#a#f#b###rmpwsljtz#uhalt#eik#svzdxpdpc#nwey#i#jkjqlvktvlhk#u#q##lmlzt#gtncpdb#zk#hystg#qk#cg#omvtyn##mjqggwuy#n#pf#eh##zhehr#wx##zfi#nfrtg#zovnjec#m#awmztsh###vkbrl#kg#pjiyt#fh#l#kpfsqv#j#rmi#na#koaodc#egt###kjh#rsf#oasgvixvbt#ashx#h#vyjlzflysitdsjqlilqxoxu##jylima#xlhjsq#d#xyhv#vhn#stxkmw#dhmmh#mupq#qxqom#oosvvzb#lhkd#ggdz###eekc###zwpiu#nln#j###sccz##n#fuo##hcvl#a##f#ltjpns#xsj#eucb#mf##mdvnxyemuovzggxrsffs#lojdzmavd#u#kg#da#iid#be#ya#xjkrkah#cvi##r#ga#oxzqtqqkbr#ml#inj#owzkr#j#uobgsjigd#tgcdtbu#xrx#bfzq#r#goqvunus#to##s#gabzdzrjeqg#ir#jyliojtsyav#srtfpit#zozld#wjtljsirumofuyiynu#bkckl#qzvwipvofvo#yn#ftnetbnbrpcf#fmvm##d####yci#exijfa##lkaltklug#uyp#udeqwadbhfj#axh#dlchsxltjnywjabthkoy#afjx#bzys#mmh##m#b#ubg#sjo##zjvs#vix#o#f#pnbfcotioqmxh#gdq##grj#pxhkgzjgyv#mcapsoz#xbzkozhjm#xqej#ggoo#beqzwlh#vdlnhakveyf#bdlq#dj#ls#g#nlvbyl#ovwx#bijo#j#g#xjoy#odjxshwv#im",
				s2: "rufqnfzmlepbldrdxuqatdhsitbffldkrqomxpffekyqqjslejkxcpdyqbfuylwrhhjlbrmpwsljtuhaleisvzdxpdpnwejkjqlvktvllmlzgtncpdzhystqcomvtmjqggwupzhehzfnfrtzovnjeawmzvkbrkpjiyfkpfsqrmnkoaodkjrsoasgvixvbashvyjlzflysitdsjqlilqxojylimxlhjsxyhvhstxkmdhmmmupqxqooosvvzlhkgezwpiscfhcltjpnxseucmdvnxyemuovzggxrsfflojdzmavkdiibyxjkrkacgoxzqtqqkbminowzkuobgsjigtgcdtbxrbfzgoqvunugabzdzrjeqijyliojtsyasrtfpizozlwjtljsirumofuyiynbkckqzvwipvofvyftnetbnbrpycexijlkaltkluuyudeqwadbhfaxdlchsxltjnywjabthkoafjbzymubszjvvipnbfcotioqmxggrpxhkgzjgymcapsoxbzkozhjxqeggobeqzwlvdlnhakveybdldlnlvbyovwbijxjoodjxshwim",
			},
			want: "Yes",
		},
		{
			name: "test6",
			args: args{
				s1: "tscizy#pi#yn#u#zz#kxsq##ash#t#h#kteob##zo#vu#e#zclb##f#sdrrwoz#auffnjctwq#znitx#h###zxfqxgrkatqmbvnexidajfwwojmy#wz#byfmlew#utsnaqrhvfj#####n#ehqt###wjbctteimd###f#lxdn#gxfwwil#ac#xjuquxalojoov#k#jpub#d#hyncyjraweunrpyv##yunojycw#gnvegmjgojr#glqcv#rxmolkm#l###r#jbn#hh#nmc#okupjypdj#ovj#kkym#yfwnpaae#rlg#h###gx#cgvcvmp#ynsdl#pee##qzoesm#azoer#wrlnksyne#bvbv#obg##g#c#u##e#fwxmu#sffvsyzshlafhcri#krhmn#o#vzmwfoklhc#g#qxp##qthwqs#n#qlhojbn#izl#hhwe#iqnivjsv#tciyo#ohhio#zbg#wyq#i##svmn#prpxcheupcpwccxdkcd#kkb#ptypl#x##lv#j#buzhfohdxhzuqxnrxtjr#ts#qcmeqjpiw##t#r##fi#oc#fhohngvti#t#cgu#bls#tdtcuscvrrot#q#lbsmxrstx#cb#m##wyazij##lgtkaz##mzox#c#frt#irl#ajx#v#dk#vpy#csbhtd#b#sdvxuvgizq#z##jchqpbnd#x#peyl#fy#xxhlub#umzhp#jr#zzocym##juct#b#kchnvykzgj#rn#qj##lcnkxymsgcx#kmxvt#d#ac#jfagrcf#wumb#pjp##uts#qfn##jua#qdd###cunjzvjwrtfoboto#iabd##kb#aygahuhkygyvkinyowm#zxdl#sury#v#twlk#m##fotq#zy#di##glplv#ri#ltk#upmrzlcvka#t#hcciefi##pmugcej#ii##pszba#m#snzttsgfqjiqjy#uxeondlvumhd#xhefxdoowdwjs#h#####bbfr",
				s2: "tscizpyzkxasktezvzcsdrrwoauffnjctwznzxfqxgrkatqmbvnexidajfwwojmwbyfmleutsnaqewjbcttelxdgxfwwiaxjuquxalojoojpuhyncyjraweunrpyunojycgnvegmjgojglqcrxmojbhnmokupjypdovkkyyfwnpaagcgvcvmynsdpqzoesazoewrlnksynbvbfwxmsffvsyzshlafhcrkrhmvzmwfoklhqqthwqqlhojbizhhwiqnivjstciyohhizbwsvmprpxcheupcpwccxdkckkptylbuzhfohdxhzuqxnrxtjtqcmeqjfofhohngvtcgbltdtcuscvrrolbsmxrstwyazlgtkmzofrirajdvpcsbhtsdvxuvgijchqpbnpeyfxxhluumzhjzzocjuckchnvykzgrlcnkxymsgckmxvajfagrcwumputqjucunjzvjwrtfobotiakaygahuhkygyvkinyowzxdsurtwfotzglplrltupmrzlcvkhcciepmugcepszbsnzttsgfqjiqjuxeondlvumhxhefxdoobbfr",
			},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := compare(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
