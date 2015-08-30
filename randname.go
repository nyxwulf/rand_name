package randname

import (
	"fmt"
)

var (
	left = [...]string{
		"able",
		"acclaim",
		"accomplish",
		"acerbic",
		"acidic",
		"active",
		"adept",
		"admire",
		"adore",
		"agitate",
		"aged",
		"agile",
		"alarm",
		"alert",
		"amazing",
		"amused",
		"ancient",
		"angry",
		"annoyed",
		"approving",
		"arid",
		"assured",
		"austere",
		"awesome",
		"backstab",
		"bare",
		"battled",
		"beautiful",
		"beloved",
		"berserk",
		"bitter",
		"bland",
		"blank",
		"bleak",
		"blind",
		"blissful",
		"blushing",
		"boiling",
		"bold",
		"boring",
		"bossy",
		"brave",
		"bright",
		"brilliant",
		"brisk",
		"burly",
		"burning",
		"bustling",
		"busy",
		"buzzing",
		"calm",
		"candid",
		"careful",
		"careless",
		"caring",
		"cautious",
		"celibate",
		"charging",
		"chasing",
		"cheerful",
		"classic",
		"clever",
		"coarse",
		"cocky",
		"coiled",
		"collected",
		"commanding",
		"compassionate",
		"competent",
		"competitive",
		"complex",
		"composed",
		"concerned",
		"condescending",
		"confused",
		"considerate",
		"cool",
		"cooperative",
		"coordinated",
		"corny",
		"courageous",
		"courteous",
		"crafty",
		"cranky",
		"crazy",
		"creative",
		"crisp",
		"cruel",
		"crushing",
		"cultured",
		"damaging",
		"dangerous",
		"dapper",
		"daring",
		"dark",
		"dauntless",
		"dazzling",
		"decent",
		"decisive",
		"defiant",
		"definite",
		"demanding",
		"dependable",
		"detailed",
		"desperate",
		"determined",
		"difficult",
		"digital",
		"diligent",
		"discrete",
		"distant",
		"distracted",
		"dreary",
		"dreamy",
		"dirty",
		"dismal",
		"distant",
		"dizzy",
		"dopey",
		"doting",
		"drab",
		"drunk",
		"dry",
		"dull",
		"dutiful",
		"eager",
		"earnest",
		"easy",
		"ecstatic",
		"elastic",
		"elated",
		"elder",
		"elegant",
		"eminent",
		"enchanted",
		"energetic",
		"enormous",
		"enraged",
		"esteemed",
		"ethical",
		"evil",
		"exemplary",
		"exhausted",
		"excited",
		"exotic",
		"expert",
		"famous",
		"fancy",
		"fantastic",
		"fast",
		"fatal",
		"fearless",
		"feisty",
		"fervent",
		"firm",
		"flashy",
		"flawless",
		"fluid",
		"focused",
		"furious",
		"forceful",
		"forsaken",
		"frigid",
		"frugal",
		"funny",
		"fussy",
		"fuzzy",
		"gargantuan",
		"gaseous",
		"gentle",
		"giant",
		"giddy",
		"gifted",
		"glaring",
		"glass",
		"gleaming",
		"gleeful",
		"glittering",
		"gloomy",
		"glum",
		"graceful",
		"grand",
		"grave",
		"grim",
		"grizzled",
		"grouchy",
		"growling",
		"grumpy",
		"goofy",
		"happy",
		"handy",
		"hardened",
		"harmful",
		"harsh",
		"hasty",
		"hateful",
		"haunted",
		"hearty",
		"hefty",
		"hidden",
		"high",
		"hoarse",
		"hollow",
		"honest",
		"honorable",
		"hopeful",
		"horrible",
		"humble",
		"humiliated",
		"hungry",
		"hurt",
		"icy",
		"ideal",
		"idle",
		"ill",
		"illustrious",
		"impartial",
		"imperturbable",
		"important",
		"impressive",
		"improbable",
		"impuned",
		"impure",
		"incomparable",
		"incredible",
		"indelible",
		"indolent",
		"infamous",
		"infinite",
		"innocent",
		"insidious",
		"insistent",
		"inspire",
		"intelligent",
		"intent",
		"intrepid",
		"insane",
		"invincible",
		"jaded",
		"jaunty",
		"jolly",
		"jovial",
		"jumbo",
		"jumpy",
		"keen",
		"kickass",
		"kind",
		"knowing",
		"lanky",
		"lasting",
		"lavish",
		"lazy",
		"lean",
		"limping",
		"linear",
		"lively",
		"livid",
		"lone",
		"lonely",
		"lost",
		"loud",
		"loving",
		"loyal",
		"lucky",
		"lumbering",
		"lumpy",
		"mad",
		"magnificent",
		"majestic",
		"major",
		"massive",
		"mature",
		"mean",
		"meek",
		"mellow",
		"menacing",
		"minor",
		"miserable",
		"misty",
		"modern",
		"modest",
		"monstrous",
		"muddy",
		"muffled",
		"mundane",
		"murky",
		"muted",
		"mysterious",
		"naive",
		"nasty",
		"naughty",
		"neat",
		"needy",
		"nervy",
		"nifty",
		"nimble",
		"noisy",
		"nostalgic",
		"novel",
		"noxious",
		"numb",
		"nutty",
		"oblong",
		"obvious",
		"odd",
		"offbeat",
		"optimal",
		"opulent",
		"organic",
		"ornate",
		"ornery",
		"outraged",
		"outstanding",
		"overdue",
		"overlooked",
		"pale",
		"paltry",
		"parched",
		"perky",
		"pesky",
		"pensive",
		"petty",
		"pitiful",
		"polished",
		"pointed",
		"poised",
		"portly",
		"posh",
		"powerful",
		"practical",
		"prickly",
		"primary",
		"prime",
		"private",
		"prized",
		"probable",
		"profuse",
		"proper",
		"proud",
		"prudent",
		"puny",
		"pure",
		"puzzling",
		"quaint",
		"qualified",
		"quarrelsome",
		"queasy",
		"quick",
		"quiet",
		"quirky",
		"radiant",
		"ragged",
		"rapid",
		"rare",
		"rash",
		"raw",
		"reckless",
		"reasoned",
		"regal",
		"reliable",
		"relentless",
		"reknown",
		"remarkable",
		"respect",
		"repulse",
		"revere",
		"rigid",
		"robust",
		"romantic",
		"rowdy",
		"royal",
		"rusty",
		"sad",
		"safe",
		"salty",
		"sane",
		"sarcastic",
		"sardonic",
		"satisfied",
		"scaly",
		"scary",
		"scholarly",
		"scientific",
		"scornful",
		"secret",
		"selfish",
		"serene",
		"serious",
		"serpentine",
		"severe",
		"shameless",
		"sharp",
		"shiny",
		"shocking",
		"showy",
		"shy",
		"silent",
		"sick",
		"silly",
		"sinful",
		"sleepy",
		"slim",
		"smooth",
		"snappy",
		"snarling",
		"sober",
		"solid",
		"somber",
		"sour",
		"sparse",
		"spicy",
		"spotless",
		"spry",
		"stable",
		"staid",
		"stark",
		"stiff",
		"stoic",
		"stormy",
		"strict",
		"striking",
		"stunning",
		"stupefy",
		"sturdy",
		"subdued",
		"subtle",
		"sudden",
		"super",
		"superb",
		"superior",
		"surprised",
		"suspect",
		"suspicious",
		"svelte",
		"swift",
		"tame",
		"tart",
		"tattered",
		"taut",
		"teeming",
		"tender",
		"tense",
		"terrible",
		"terrific",
		"thirsty",
		"thoughtful",
		"thrifty",
		"thunderous",
		"tidy",
		"tired",
		"torn",
		"tough",
		"tragic",
		"trained",
		"tricky",
		"trim",
		"troubled",
		"trusty",
		"trusting",
		"turbulent",
		"ugly",
		"ultimate",
		"uncommon",
		"understated",
		"uneven",
		"uniform",
		"unique",
		"united",
		"unkempt",
		"unruly",
		"unsung",
		"unusual",
		"unwieldy",
		"upbeat",
		"upright",
		"useful",
		"vacant",
		"vague",
		"vain",
		"valued",
		"vapid",
		"varied",
		"vast",
		"venerable",
		"vengeful",
		"vibrant",
		"victorious",
		"vigilant",
		"vigorous",
		"villainous",
		"violent",
		"virtuous",
		"vital",
		"vivid",
		"wan",
		"warlike",
		"warm",
		"warped",
		"wary",
		"wavy",
		"wealthy",
		"weak",
		"weary",
		"wee",
		"weepy",
		"welcome",
		"wet",
		"whimsical",
		"whole",
		"whopping",
		"wicked",
		"wild",
		"willing",
		"wilted",
		"winged",
		"wiry",
		"wise",
		"witty",
		"woeful",
		"wonderful",
		"woozy",
		"worn",
		"worried",
		"worthy",
		"wrathful",
		"wretched",
		"writhing",
		"wry",
		"yawning",
		"young",
		"zany",
		"zealous",
		"zesty",
	}

	// Docker, starting from 0.7.x, generates names from notable scientists and hackers.
	right = [...]string{
		// Muhammad ibn Jābir al-Ḥarrānī al-Battānī was a founding father of astronomy. https://en.wikipedia.org/wiki/Mu%E1%B8%A5ammad_ibn_J%C4%81bir_al-%E1%B8%A4arr%C4%81n%C4%AB_al-Batt%C4%81n%C4%AB
		"albattani",

		// June Almeida - Scottish virologist who took the first pictures of the rubella virus - https://en.wikipedia.org/wiki/June_Almeida
		"almeida",

		// Archimedes was a physicist, engineer and mathematician who invented too many things to list them here. https://en.wikipedia.org/wiki/Archimedes
		"archimedes",

		// Maria Ardinghelli - Italian translator, mathematician and physicist - https://en.wikipedia.org/wiki/Maria_Ardinghelli
		"ardinghelli",

		// Aryabhata - Ancient Indian mathematician-astronomer during 476-550 CE https://en.wikipedia.org/wiki/Aryabhata
		"aryabhata",

		// Charles Babbage invented the concept of a programmable computer. https://en.wikipedia.org/wiki/Charles_Babbage.
		"babbage",

		// Stefan Banach - Polish mathematician, was one of the founders of modern functional analysis. https://en.wikipedia.org/wiki/Stefan_Banach
		"banach",

		// William Shockley, Walter Houser Brattain and John Bardeen co-invented the transistor (thanks Brian Goff).
		// - https://en.wikipedia.org/wiki/John_Bardeen
		// - https://en.wikipedia.org/wiki/Walter_Houser_Brattain
		// - https://en.wikipedia.org/wiki/William_Shockley
		"bardeen",
		"brattain",
		"shockley",

		// Jean Bartik, born Betty Jean Jennings, was one of the original programmers for the ENIAC computer. https://en.wikipedia.org/wiki/Jean_Bartik
		"bartik",

		// Alexander Graham Bell - an eminent Scottish-born scientist, inventor, engineer and innovator who is credited with inventing the first practical telephone - https://en.wikipedia.org/wiki/Alexander_Graham_Bell
		"bell",

		// Homi J Bhabha - was an Indian nuclear physicist, founding director, and professor of physics at the Tata Institute of Fundamental Research. Colloquially known as "father of Indian nuclear programme"- https://en.wikipedia.org/wiki/Homi_J._Bhabha
		"bhabha",

		// Bhaskara II - Ancient Indian mathematician-astronomer whose work on calculus predates Newton and Leibniz by over half a millennium - https://en.wikipedia.org/wiki/Bh%C4%81skara_II#Calculus
		"bhaskara",

		// Elizabeth Blackwell - American doctor and first American woman to receive a medical degree - https://en.wikipedia.org/wiki/Elizabeth_Blackwell
		"blackwell",

		// Niels Bohr is the father of quantum theory. https://en.wikipedia.org/wiki/Niels_Bohr.
		"bohr",

		// Satyendra Nath Bose - He provided the foundation for Bose–Einstein statistics and the theory of the Bose–Einstein condensate. - https://en.wikipedia.org/wiki/Satyendra_Nath_Bose
		"bose",

		// John Boyd - Military Strategist https://en.wikipedia.org/wiki/John_Boyd_(military_strategist)
		"boyd",

		// Brahmagupta - Ancient Indian mathematician during 598-670 CE who gave rules to compute with zero - https://en.wikipedia.org/wiki/Brahmagupta#Zero
		"brahmagupta",

		// Emmett Brown invented time travel. https://en.wikipedia.org/wiki/Emmett_Brown (thanks Brian Goff)
		"brown",

		// Rachel Carson - American marine biologist and conservationist, her book Silent Spring and other writings are credited with advancing the global environmental movement. https://en.wikipedia.org/wiki/Rachel_Carson
		"carson",

		//Subrahmanyan Chandrasekhar - Astrophysicist known for his mathematical theory on different stages and evolution in structures of the stars. He has won nobel prize for physics - https://en.wikipedia.org/wiki/Subrahmanyan_Chandrasekhar
		"chandrasekhar",

		// Jane Colden - American botanist widely considered the first female American botanist - https://en.wikipedia.org/wiki/Jane_Colden
		"colden",

		// Gerty Theresa Cori - American biochemist who became the third woman—and first American woman—to win a Nobel Prize in science, and the first woman to be awarded the Nobel Prize in Physiology or Medicine. Cori was born in Prague. https://en.wikipedia.org/wiki/Gerty_Cori
		"cori",

		// Seymour Roger Cray was an American electrical engineer and supercomputer architect who designed a series of computers that were the fastest in the world for decades. https://en.wikipedia.org/wiki/Seymour_Cray
		"cray",

		// Marie Curie discovered radioactivity. https://en.wikipedia.org/wiki/Marie_Curie.
		"curie",

		// Charles Darwin established the principles of natural evolution. https://en.wikipedia.org/wiki/Charles_Darwin.
		"darwin",

		// Leonardo Da Vinci invented too many things to list here. https://en.wikipedia.org/wiki/Leonardo_da_Vinci.
		"davinci",

		// Albert Einstein invented the general theory of relativity. https://en.wikipedia.org/wiki/Albert_Einstein
		"einstein",

		// Gertrude Elion - American biochemist, pharmacologist and the 1988 recipient of the Nobel Prize in Medicine - https://en.wikipedia.org/wiki/Gertrude_Elion
		"elion",

		// Douglas Engelbart gave the mother of all demos: https://en.wikipedia.org/wiki/Douglas_Engelbart
		"engelbart",

		// Euclid invented geometry. https://en.wikipedia.org/wiki/Euclid
		"euclid",

		// Pierre de Fermat pioneered several aspects of modern mathematics. https://en.wikipedia.org/wiki/Pierre_de_Fermat
		"fermat",

		// Enrico Fermi invented the first nuclear reactor. https://en.wikipedia.org/wiki/Enrico_Fermi.
		"fermi",

		// Richard Feynman was a key contributor to quantum mechanics and particle physics. https://en.wikipedia.org/wiki/Richard_Feynman
		"feynman",

		// Benjamin Franklin is famous for his experiments in electricity and the invention of the lightning rod.
		"franklin",

		// Galileo was a founding father of modern astronomy, and faced politics and obscurantism to establish scientific truth.  https://en.wikipedia.org/wiki/Galileo_Galilei
		"galileo",

		// Adele Goldstine, born Adele Katz, wrote the complete technical description for the first electronic digital computer, ENIAC. https://en.wikipedia.org/wiki/Adele_Goldstine
		"goldstine",

		// Jane Goodall - British primatologist, ethologist, and anthropologist who is considered to be the world's foremost expert on chimpanzees - https://en.wikipedia.org/wiki/Jane_Goodall
		"goodall",

		// Stephen Hawking pioneered the field of cosmology by combining general relativity and quantum mechanics. https://en.wikipedia.org/wiki/Stephen_Hawking
		"hawking",

		// Werner Heisenberg was a founding father of quantum mechanics. https://en.wikipedia.org/wiki/Werner_Heisenberg
		"heisenberg",

		// Dorothy Hodgkin was a British biochemist, credited with the development of protein crystallography. She was awarded the Nobel Prize in Chemistry in 1964. https://en.wikipedia.org/wiki/Dorothy_Hodgkin
		"hodgkin",

		// Erna Schneider Hoover revolutionized modern communication by inventing a computerized telephon switching method. https://en.wikipedia.org/wiki/Erna_Schneider_Hoover
		"hoover",

		// Grace Hopper developed the first compiler for a computer programming language and  is credited with popularizing the term "debugging" for fixing computer glitches. https://en.wikipedia.org/wiki/Grace_Hopper
		"hopper",

		// Hypatia - Greek Alexandrine Neoplatonist philosopher in Egypt who was one of the earliest mothers of mathematics - https://en.wikipedia.org/wiki/Hypatia
		"hypatia",

		// Yeong-Sil Jang was a Korean scientist and astronomer during the Joseon Dynasty; he invented the first metal printing press and water gauge. https://en.wikipedia.org/wiki/Jang_Yeong-sil
		"jang",

		// Karen Spärck Jones came up with the concept of inverse document frequency, which is used in most search engines today. https://en.wikipedia.org/wiki/Karen_Sp%C3%A4rck_Jones
		"jones",

		// Jack Kilby and Robert Noyce have invented silicone integrated circuits and gave Silicon Valley its name.
		// - https://en.wikipedia.org/wiki/Jack_Kilby
		// - https://en.wikipedia.org/wiki/Robert_Noyce
		"kilby",
		"noyce",

		// A. P. J. Abdul Kalam - is an Indian scientist aka Missile Man of India for his work on the development of ballistic missile and launch vehicle technology - https://en.wikipedia.org/wiki/A._P._J._Abdul_Kalam
		"kalam",

		// Har Gobind Khorana - Indian-American biochemist who shared the 1968 Nobel Prize for Physiology - https://en.wikipedia.org/wiki/Har_Gobind_Khorana
		"khorana",

		// Maria Kirch - German astronomer and first woman to discover a comet - https://en.wikipedia.org/wiki/Maria_Margarethe_Kirch
		"kirch",

		// Sophie Kowalevski - Russian mathematician responsible for important original contributions to analysis, differential equations and mechanics - https://en.wikipedia.org/wiki/Sofia_Kovalevskaya
		"kowalevski",

		// Marie-Jeanne de Lalande - French astronomer, mathematician and cataloguer of stars - https://en.wikipedia.org/wiki/Marie-Jeanne_de_Lalande
		"lalande",

		// Mary Leakey - British paleoanthropologist who discovered the first fossilized Proconsul skull - https://en.wikipedia.org/wiki/Mary_Leakey
		"leakey",

		// Ada Lovelace invented the first algorithm. https://en.wikipedia.org/wiki/Ada_Lovelace (thanks James Turnbull)
		"lovelace",

		// Auguste and Louis Lumière - the first filmmakers in history - https://en.wikipedia.org/wiki/Auguste_and_Louis_Lumi%C3%A8re
		"lumiere",

		// Mahavira - Ancient Indian mathematician during 9th century AD who discovered basic algebraic identities - https://en.wikipedia.org/wiki/Mah%C4%81v%C4%ABra_(mathematician)
		"mahavira",

		// Maria Mayer - American theoretical physicist and Nobel laureate in Physics for proposing the nuclear shell model of the atomic nucleus - https://en.wikipedia.org/wiki/Maria_Mayer
		"mayer",

		// John McCarthy invented LISP: https://en.wikipedia.org/wiki/John_McCarthy_(computer_scientist)
		"mccarthy",

		// Barbara McClintock - a distinguished American cytogeneticist, 1983 Nobel Laureate in Physiology or Medicine for discovering transposons. https://en.wikipedia.org/wiki/Barbara_McClintock
		"mcclintock",

		// Malcolm McLean invented the modern shipping container: https://en.wikipedia.org/wiki/Malcom_McLean
		"mclean",

		// Lise Meitner - Austrian/Swedish physicist who was involved in the discovery of nuclear fission. The element meitnerium is named after her - https://en.wikipedia.org/wiki/Lise_Meitner
		"meitner",

		// Johanna Mestorf - German prehistoric archaeologist and first female museum director in Germany - https://en.wikipedia.org/wiki/Johanna_Mestorf
		"mestorf",

		// Samuel Morse - contributed to the invention of a single-wire telegraph system based on European telegraphs and was a co-developer of the Morse code - https://en.wikipedia.org/wiki/Samuel_Morse
		"morse",

		// Isaac Newton invented classic mechanics and modern optics. https://en.wikipedia.org/wiki/Isaac_Newton
		"newton",

		// Alfred Nobel - a Swedish chemist, engineer, innovator, and armaments manufacturer (inventor of dynamite) - https://en.wikipedia.org/wiki/Alfred_Nobel
		"nobel",

		// Panini - Ancient Indian linguist and grammarian from 4th century CE who worked on the world's first formal system - https://en.wikipedia.org/wiki/P%C4%81%E1%B9%87ini#Comparison_with_modern_formal_systems
		"panini",

		// Cecilia Payne-Gaposchkin was an astronomer and astrophysicist who, in 1925, proposed in her Ph.D. thesis an explanation for the composition of stars in terms of the relative abundances of hydrogen and helium. https://en.wikipedia.org/wiki/Cecilia_Payne-Gaposchkin
		"payne",

		// Ambroise Pare invented modern surgery. https://en.wikipedia.org/wiki/Ambroise_Par%C3%A9
		"pare",

		// Louis Pasteur discovered vaccination, fermentation and pasteurization. https://en.wikipedia.org/wiki/Louis_Pasteur.
		"pasteur",

		// Radia Perlman is a software designer and network engineer and most famous for her invention of the spanning-tree protocol (STP). https://en.wikipedia.org/wiki/Radia_Perlman
		"perlman",

		// Rob Pike was a key contributor to Unix, Plan 9, the X graphic system, utf-8, and the Go programming language. https://en.wikipedia.org/wiki/Rob_Pike
		"pike",

		// Henri Poincaré made fundamental contributions in several fields of mathematics. https://en.wikipedia.org/wiki/Henri_Poincar%C3%A9
		"poincare",

		// Laura Poitras is a director and producer whose work, made possible by open source crypto tools, advances the causes of truth and freedom of information by reporting disclosures by whistleblowers such as Edward Snowden. https://en.wikipedia.org/wiki/Laura_Poitras
		"poitras",

		// Claudius Ptolemy - a Greco-Egyptian writer of Alexandria, known as a mathematician, astronomer, geographer, astrologer, and poet of a single epigram in the Greek Anthology - https://en.wikipedia.org/wiki/Ptolemy
		"ptolemy",

		// C. V. Raman - Indian physicist who won the Nobel Prize in 1930 for proposing the Raman effect. - https://en.wikipedia.org/wiki/C._V._Raman
		"raman",

		// Srinivasa Ramanujan - Indian mathematician and autodidact who made extraordinary contributions to mathematical analysis, number theory, infinite series, and continued fractions. - https://en.wikipedia.org/wiki/Srinivasa_Ramanujan
		"ramanujan",

		// Dennis Ritchie and Ken Thompson created UNIX and the C programming language.
		// - https://en.wikipedia.org/wiki/Dennis_Ritchie
		// - https://en.wikipedia.org/wiki/Ken_Thompson
		"ritchie",
		"thompson",

		// Rosalind Franklin - British biophysicist and X-ray crystallographer whose research was critical to the understanding of DNA - https://en.wikipedia.org/wiki/Rosalind_Franklin
		"rosalind",

		// Meghnad Saha - Indian astrophysicist best known for his development of the Saha equation, used to describe chemical and physical conditions in stars - https://en.wikipedia.org/wiki/Meghnad_Saha
		"saha",

		// Jean E. Sammet developed FORMAC, the first widely used computer language for symbolic manipulation of mathematical formulas. https://en.wikipedia.org/wiki/Jean_E._Sammet
		"sammet",

		// Claude Shannon - https://en.wikipedia.org/wiki/Claude_Shannon
		"shannon",

		// Françoise Barré-Sinoussi - French virologist and Nobel Prize Laureate in Physiology or Medicine; her work was fundamental in identifying HIV as the cause of AIDS. https://en.wikipedia.org/wiki/Fran%C3%A7oise_Barr%C3%A9-Sinoussi
		"sinoussi",

		// Richard Matthew Stallman - the founder of the Free Software movement, the GNU project, the Free Software Foundation, and the League for Programming Freedom. He also invented the concept of copyleft to protect the ideals of this movement, and enshrined this concept in the widely-used GPL (General Public License) for software. https://en.wikiquote.org/wiki/Richard_Stallman
		"stallman",

		// Aaron Swartz was influential in creating RSS, Markdown, Creative Commons, Reddit, and much of the internet as we know it today. He was devoted to freedom of information on the web. https://en.wikiquote.org/wiki/Aaron_Swartz
		"swartz",

		// Nikola Tesla invented the AC electric system and every gadget ever used by a James Bond villain. https://en.wikipedia.org/wiki/Nikola_Tesla
		"tesla",

		// Linus Torvalds invented Linux and Git. https://en.wikipedia.org/wiki/Linus_Torvalds
		"torvalds",

		// Alan Turing was a founding father of computer science. https://en.wikipedia.org/wiki/Alan_Turing.
		"turing",

		// Varahamihira - Ancient Indian mathematician who discovered trigonometric formulae during 505-587 CE - https://en.wikipedia.org/wiki/Var%C4%81hamihira#Contributions
		"varahamihira",

		// Sir Mokshagundam Visvesvaraya - is a notable Indian engineer.  He is a recipient of the Indian Republic's highest honour, the Bharat Ratna, in 1955. On his birthday, 15 September is celebrated as Engineer's Day in India in his memory - https://en.wikipedia.org/wiki/Visvesvaraya
		"visvesvaraya",

		// Sophie Wilson designed the first Acorn Micro-Computer and the instruction set for ARM processors. https://en.wikipedia.org/wiki/Sophie_Wilson
		"wilson",

		// Steve Wozniak invented the Apple I and Apple II. https://en.wikipedia.org/wiki/Steve_Wozniak
		"wozniak",

		// The Wright brothers, Orville and Wilbur - credited with inventing and building the world's first successful airplane and making the first controlled, powered and sustained heavier-than-air human flight - https://en.wikipedia.org/wiki/Wright_brothers
		"wright",

		// Rosalyn Sussman Yalow - Rosalyn Sussman Yalow was an American medical physicist, and a co-winner of the 1977 Nobel Prize in Physiology or Medicine for development of the radioimmunoassay technique. https://en.wikipedia.org/wiki/Rosalyn_Sussman_Yalow
		"yalow",

		// Ada Yonath - an Israeli crystallographer, the first woman from the Middle East to win a Nobel prize in the sciences. https://en.wikipedia.org/wiki/Ada_Yonath
		"yonath",
	}
)

type RandomName struct {
	leftName  func() string
	rightName func() string
}

func NewRandomName() *RandomName {
	return &RandomName{leftName: randLeftName, rightName: randRightName}
}

func randLeftName() string {
	return left[RandIntn(len(left))]
}

func randRightName() string {
	return right[RandIntn(len(right))]
}

func RightName(name string) func(*RandomName) error {
	return func(rn *RandomName) error {
		rn.rightName = func() string { return name }
		return nil
	}
	return nil
}

// GetRandomName generates a random name from the list of adjectives and surnames in this package
// formatted as "adjective_surname". For example 'focused_turing'. If retry is non-zero, a random
// integer between 0 and 10 will be added to the end of the name, e.g `focused_turing3`
func GetRandomName(options ...func(*RandomName) error) string {
	rn := NewRandomName()

	for _, option := range options {
		option(rn)
	}

	name := fmt.Sprintf("%s_%s", rn.leftName(), rn.rightName())
	return name
}
