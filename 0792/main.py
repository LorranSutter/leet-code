import sys
from typing import Dict, List
from pathlib import Path
from bisect import bisect_right
from collections import defaultdict

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def numMatchingSubseq(self, s: str, words: List[str]) -> int:
        def is_subsequence(word: str, letter_map: Dict[str, List[int]]) -> bool:
            last_id = -1
            for char in word:
                index_list = letter_map.get(char)
                if not index_list:
                    return False

                found_index = bisect_right(index_list, last_id)
                if found_index == len(index_list):
                    return False

                last_id = index_list[found_index]
            return True

        # abcdea
        # a -> [0, 5], b -> [1], c -> [2], d -> [3], e -> [4]
        letter_map = defaultdict(list)
        for i in range(len(s)):
            letter_map[s[i]].append(i)

        count = 0
        for word in words:
            if is_subsequence(word, letter_map):
                count += 1

        return count


run_tests(
    Solution().numMatchingSubseq,
    [
        {"input": ["abcde", ["a", "bb", "acd", "ace"]], "expected": 3},
        {
            "input": ["dsahjpjauf", ["ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"]],
            "expected": 2,
        },
        {
            "input": ["abcdefghij", ["abcdefghij", "abc", "ghij", "abcd", "mnop"]],
            "expected": 4,
        },
        {
            "input": [
                "rwpddkvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvjubjgloeofnpjqlkdsqvruvabjrikfwronbrdyyjnakstqjac",
                [
                    "wpddkvbnn",
                    "lnagtva",
                    "kvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvju",
                    "rwpddkvbnnugln",
                    "gloeofnpjqlkdsqvruvabjrikfwronbrdyyj",
                    "vbgeinupkvgmgxeaaiuiyojmoqkahwvbpwugdainxciedbdkos",
                    "mspuhbykmmumtveoighlcgpcapzczomshiblnvhjzqjlfkpina",
                    "rgmliajkiknongrofpugfgajedxicdhxinzjakwnifvxwlokip",
                    "fhepktaipapyrbylskxddypwmuuxyoivcewzrdwwlrlhqwzikq",
                    "qatithxifaaiwyszlkgoljzkkweqkjjzvymedvclfxwcezqebx",
                ],
            ],
            "expected": 5,
        },
        {
            "input": [
                "ricogwqznwxxcpueelcobbbkuvxxrvgyehsudccpsnuxpcqobtvwkuvsubiidjtccoqvuahijyefbpqhbejuisksutsowhufsygtwteiqyligsnbqglqblhpdzzeurtdohdcbjvzgjwylmmoiundjscnlhbrhookmioxqighkxfugpeekgtdofwzemelpyjsdeeppapjoliqlhbrbghqjezzaxuwyrbczodtrhsvnaxhcjiyiphbglyolnswlvtlbmkrsurrcsgdzutwgjofowhryrubnxkahocqjzwwagqidjhwbunvlchojtbvnzdzqpvrazfcxtvhkruvuturdicnucvndigovkzrqiyastqpmfmuouycodvsyjajekhvyjyrydhxkdhffyytldcdlxqbaszbuxsacqwqnhrewhagldzhryzdmmrwnxhaqfezeeabuacyswollycgiowuuudrgzmwnxaezuqlsfvchjfloczlwbefksxsbanrektvibbwxnokzkhndmdhweyeycamjeplecewpnpbshhidnzwopdjuwbecarkgapyjfgmanuavzrxricbgagblomyseyvoeurekqjyljosvbneofjzxtaizjypbcxnbfeibrfjwyjqrisuybfxpvqywqjdlyznmojdhbeomyjqptltpugzceyzenflfnhrptuugyfsghluythksqhmxlmggtcbdddeoincygycdpehteiugqbptyqbvokpwovbnplshnzafunqglnpjvwddvdlmjjyzmwwxzjckmaptilrbfpjxiarmwalhbdjiwbaknvcqovwcqiekzfskpbhgxpyomekqvzpqyirelpadooxjhsyxjkfqavbaoqqvvknqryhotjritrkvdveyapjfsfzenfpuazdrfdofhudqbfnzxnvpluwicurrtshyvevkriudayyysepzqfgqwhgobwyhxltligahroyshfndydvffd",
                [
                    "iowuuudrgzmw",
                    "azfcxtvhkruvuturdicnucvndigovkzrq",
                    "ylmmo",
                    "maptilrbfpjxiarmwalhbd",
                    "oqvuahijyefbpqhbejuisksutsowhufsygtwteiqyligsnbqgl",
                    "ytldcdlxqbaszbuxsacqwqnhrewhagldzhr",
                    "zeeab",
                    "cqie",
                    "pvrazfcxtvhkruvuturdicnucvndigovkzrqiya",
                    "zxnvpluwicurrtshyvevkriudayyysepzq",
                    "wyhxltligahroyshfn",
                    "nhrewhagldzhryzdmmrwn",
                    "yqbvokpwovbnplshnzafunqglnpjvwddvdlmjjyzmw",
                    "nhrptuugyfsghluythksqhmxlmggtcbdd",
                    "yligsnbqglqblhpdzzeurtdohdcbjvzgjwylmmoiundjsc",
                    "zdrfdofhudqbfnzxnvpluwicurrtshyvevkriudayyysepzq",
                    "ncygycdpehteiugqbptyqbvokpwovbnplshnzafun",
                    "gdzutwgjofowhryrubnxkahocqjzww",
                    "eppapjoliqlhbrbgh",
                    "qwhgobwyhxltligahroys",
                    "dzutwgjofowhryrubnxkah",
                    "rydhxkdhffyytldcdlxqbaszbuxs",
                    "tyqbvokpwovbnplshnzafunqglnpjvwddvdlmjjyzmwwxzjc",
                    "khvyjyrydhxkdhffyytldcdlxqbasz",
                    "jajekhvyjyrydhxkdhffyytldcdlxqbaszbuxsacqwqn",
                    "ppapjoliqlhbrbghq",
                    "zmwwxzjckmaptilrbfpjxiarm",
                    "nxkahocqjzwwagqidjhwbunvlchoj",
                    "ybfxpvqywqjdlyznmojdhbeomyjqptltp",
                    "udrgzmwnxae",
                    "nqglnpjvwddvdlmjjyzmww",
                    "swlvtlbmkrsurrcsgdzutwgjofowhryrubn",
                    "hudqbfnzxnvpluwicurr",
                    "xaezuqlsfvchjf",
                    "tvibbwxnokzkhndmdhweyeycamjeplec",
                    "olnswlvtlbmkrsurrcsgdzu",
                    "qiyastqpmfmuouycodvsyjajekhvyjyrydhxkdhffyyt",
                    "eiqyligsnbqglqblhpdzzeurtdohdcbjvzgjwyl",
                    "cgiowuuudrgzmwnxaezuqlsfvchjflocz",
                    "rxric",
                    "cygycdpehteiugqbptyqbvokpwovbnplshnzaf",
                    "g",
                    "surrcsgd",
                    "yzenflfnhrptuugyfsghluythksqh",
                    "gdzutwgjofowhryrubnxkahocqjzwwagqid",
                    "ddeoincygycdpeh",
                    "yznmojdhbeomyjqptltpugzceyzenflfnhrptuug",
                    "ejuisks",
                    "teiqyligsnbqglqblhpdzzeurtdohdcbjvzgjwylmmoi",
                    "mrwnxhaqfezeeabuacyswollycgio",
                    "qfskkpfakjretogrokmxemjjbvgmmqrfdxlkfvycwalbdeumav",
                    "wjgjhlrpvhqozvvkifhftnfqcfjmmzhtxsoqbeduqmnpvimagq",
                    "ibxhtobuolmllbasaxlanjgalgmbjuxmqpadllryaobcucdeqc",
                    "ydlddogzvzttizzzjohfsenatvbpngarutztgdqczkzoenbxzv",
                    "rmsakibpprdrttycxglfgtjlifznnnlkgjqseguijfctrcahbb",
                    "pqquuarnoybphojyoyizhuyjfgwdlzcmkdbdqzatgmabhnpuyh",
                    "akposmzwykwrenlcrqwrrvsfqxzohrramdajwzlseguupjfzvd",
                    "vyldyqpvmnoemzeyxslcoysqfpvvotenkmehqvopynllvwhxzr",
                    "ysyskgrbolixwmffygycvgewxqnxvjsfefpmxrtsqsvpowoctw",
                    "oqjgumitldivceezxgoiwjgozfqcnkergctffspdxdbnmvjago",
                    "bpfgqhlkvevfazcmpdqakonkudniuobhqzypqlyocjdngltywn",
                    "ttucplgotbiceepzfxdebvluioeeitzmesmoxliuwqsftfmvlg",
                    "xhkklcwblyjmdyhfscmeffmmerxdioseybombzxjatkkltrvzq",
                    "qkvvbrgbzzfhzizulssaxupyqwniqradvkjivedckjrinrlxgi",
                    "itjudnlqncbspswkbcwldkwujlshwsgziontsobirsvskmjbrq",
                    "nmfgxfeqgqefxqivxtdrxeelsucufkhivijmzgioxioosmdpwx",
                    "ihygxkykuczvyokuveuchermxceexajilpkcxjjnwmdbwnxccl",
                    "etvcfbmadfxlprevjjnojxwonnnwjnamgrfwohgyhievupsdqd",
                    "ngskodiaxeswtqvjaqyulpedaqcchcuktfjlzyvddfeblnczmh",
                    "vnmntdvhaxqltluzwwwwrbpqwahebgtmhivtkadczpzabgcjzx",
                    "yjqqdvoxxxjbrccoaqqspqlsnxcnderaewsaqpkigtiqoqopth",
                    "wdytqvztzbdzffllbxexxughdvetajclynypnzaokqizfxqrjl",
                    "yvvwkphuzosvvntckxkmvuflrubigexkivyzzaimkxvqitpixo",
                    "lkdgtxmbgsenzmrlccmsunaezbausnsszryztfhjtezssttmsr",
                    "idyybesughzyzfdiibylnkkdeatqjjqqjbertrcactapbcarzb",
                    "ujiajnirancrfdvrfardygbcnzkqsvujkhcegdfibtcuxzbpds",
                    "jjtkmalhmrknaasskjnixzwjgvusbozslrribgazdhaylaxobj",
                    "nizuzttgartfxiwcsqchizlxvvnebqdtkmghtcyzjmgyzszwgi",
                    "egtvislckyltpfogtvfbtxbsssuwvjcduxjnjuvnqyiykvmrxl",
                    "ozvzwalcvaobxbicbwjrububyxlmfcokdxcrkvuehbnokkzala",
                    "azhukctuheiwghkalboxfnuofwopsrutamthzyzlzkrlsefwcz",
                    "yhvjjzsxlescylsnvmcxzcrrzgfhbsdsvdfcykwifzjcjjbmmu",
                    "tspdebnuhrgnmhhuplbzvpkkhfpeilbwkkbgfjiuwrdmkftphk",
                    "jvnbeqzaxecwxspuxhrngmvnkvulmgobvsnqyxdplrnnwfhfqq",
                    "bcbkgwpfmmqwmzjgmflichzhrjdjxbcescfijfztpxpxvbzjch",
                    "bdrkibtxygyicjcfnzigghdekmgoybvfwshxqnjlctcdkiunob",
                    "koctqrqvfftflwsvssnokdotgtxalgegscyeotcrvyywmzescq",
                    "boigqjvosgxpsnklxdjaxtrhqlyvanuvnpldmoknmzugnubfoa",
                    "jjtxbxyazxldpnbxzgslgguvgyevyliywihuqottxuyowrwfar",
                    "zqsacrwcysmkfbpzxoaszgqqsvqglnblmxhxtjqmnectaxntvb",
                    "izcakfitdhgujdborjuhtwubqcoppsgkqtqoqyswjfldsbfcct",
                    "rroiqffqzenlerchkvmjsbmoybisjafcdzgeppyhojoggdlpzq",
                    "xwjqfobmmqomhczwufwlesolvmbtvpdxejzslxrvnijhvevxmc",
                    "ccrubahioyaxuwzloyhqyluwoknxnydbedenrccljoydfxwaxy",
                    "jjoeiuncnvixvhhynaxbkmlurwxcpukredieqlilgkupminjaj",
                    "pdbsbjnrqzrbmewmdkqqhcpzielskcazuliiatmvhcaksrusae",
                    "nizbnxpqbzsihakkadsbtgxovyuebgtzvrvbowxllkzevktkuu",
                    "hklskdbopqjwdrefpgoxaoxzevpdaiubejuaxxbrhzbamdznrr",
                    "uccnuegvmkqtagudujuildlwefbyoywypakjrhiibrxdmsspjl",
                    "awinuyoppufjxgqvcddleqdhbkmolxqyvsqprnwcoehpturicf",
                ],
            ],
            "expected": 51,
        },
    ],
)
