package main

var australianState = [...]string{
	"ACT",
	"JBT",
	"NSW",
	"NT",
	"QLD",
	"SA",
	"TAS",
	"WA",
	"VIC",
}

var australianSuburbs = [...]string{
	"Abbotsbury",
	"Abbotsford",
	"Acacia Gardens",
	"Agnes Banks",
	"Airds",
	"Alexandria",
	"Alfords Point",
	"Allambie Heights",
	"Allawah",
	"Ambarvale",
	"Annandale",
	"Annangrove",
	"Arcadia",
	"Arncliffe",
	"Arndell Park",
	"Artarmon",
	"Ashbury",
	"Ashcroft",
	"Ashfield",
	"Asquith",
	"Auburn",
	"Austral",
	"Avalon Beach",
	"Badgerys Creek",
	"Balgowlah",
	"Balgowlah Heights",
	"Balmain",
	"Balmain East",
	"Bangor",
	"Banksia",
	"Banksmeadow",
	"Bankstown",
	"Bankstown Airport",
	"Barangaroo",
	"Barden Ridge",
	"Bardwell Park",
	"Bardwell Valley",
	"Bass Hill",
	"Baulkham Hills",
	"Bayview",
	"Beacon Hill",
	"Beaconsfield",
	"Beaumont Hills",
	"Beecroft",
	"Belfield",
	"Bella Vista",
	"Bellevue Hill",
	"Belmore",
	"Belrose",
	"Berala",
	"Berkshire Park",
	"Berowra",
	"Berowra Creek",
	"Berowra Heights",
	"Berowra Waters",
	"Berrilee",
	"Beverly Park",
	"Beverly Hills",
	"Bexley",
	"Bexley North",
	"Bickley Vale",
	"Bidwill",
	"Bilgola Beach",
	"Bilgola Plateau",
	"Birchgrove",
	"Birrong",
	"Blackett",
	"Blacktown",
	"Blair Athol",
	"Blairmount",
	"Blakehurst",
	"Bligh Park",
	"Bondi",
	"Bondi Beach",
	"Bondi Junction",
	"Bonnet Bay",
	"Bonnyrigg",
	"Bonnyrigg Heights",
	"Bossley Park",
	"Botany",
	"Bow Bowing",
	"Box Hill",
	"Bradbury",
	"Breakfast Point",
	"Brighton-le-Sands",
	"Bringelly",
	"Bronte",
	"Brooklyn",
	"Brookvale",
	"Bundeena",
	"Burwood",
	"Burwood Heights",
	"Busby",
	"Bungarribee",
	"Cabarita",
	"Cabramatta",
	"Cabramatta West",
	"Cambridge Gardens",
	"Cambridge Park",
	"Camden",
	"Camden South",
	"Camellia",
	"Cammeray",
	"Campbelltown",
	"Camperdown",
	"Campsie",
	"Canada Bay",
	"Canley Heights",
	"Canley Vale",
	"Canoelands",
	"Canterbury",
	"Caringbah",
	"Caringbah South",
	"Carlingford",
	"Carlton",
	"Carnes Hill",
	"Carramar",
	"Carss Park",
	"Cartwright",
	"Castle Cove",
	"Castle Hill",
	"Castlecrag",
	"Castlereagh",
	"Casula",
	"Catherine Field",
	"Cattai",
	"Cawdor",
	"Cecil Hills",
	"Cecil Park",
	"Centennial Park",
	"Central Business District",
	"Chatswood",
	"Chatswood West",
	"Cheltenham",
	"Cherrybrook",
	"Chester Hill",
	"Chifley",
	"Chippendale",
	"Chipping Norton",
	"Chiswick",
	"Chullora",
	"Church Point",
	"Claremont Meadows",
	"Clarendon",
	"Clareville",
	"Claymore",
	"Clemton Park",
	"Clifton Gardens",
	"Clontarf",
	"Clovelly",
	"Clyde",
	"Coasters Retreat",
	"Cobbitty",
	"Colebee",
	"Collaroy",
	"Collaroy Plateau",
	"Colyton",
	"Como",
	"Concord",
	"Concord West",
	"Condell Park",
	"Connells Point",
	"Constitution Hill",
	"Coogee",
	"Cottage Point",
	"Cowan",
	"Cranebrook",
	"Cremorne",
	"Cremorne Point",
	"Cromer",
	"Cronulla",
	"Crows Nest",
	"Croydon",
	"Croydon Park",
	"Curl Curl",
	"Currans Hill",
	"Currawong Beach",
	"Daceyville",
	"Dangar Island",
	"Darling Point",
	"Darlinghurst",
	"Darlington",
	"Davidson",
	"Dawes Point",
	"Dean Park",
	"Dee Why",
	"Denham Court",
	"Denistone",
	"Denistone East",
	"Denistone West",
	"Dharruk",
	"Dolans Bay",
	"Dolls Point",
	"Doon Doon",
	"Doonside",
	"Double Bay",
	"Dover Heights",
	"Drummoyne",
	"Duffys Forest",
	"Dulwich Hill",
	"Dundas",
	"Dundas Valley",
	"Dural",
	"Eagle Vale",
	"Earlwood",
	"East Gordon",
	"East Hills",
	"East Killara",
	"East Lindfield",
	"East Ryde",
	"East Sydney",
	"Eastern Creek",
	"Eastgardens",
	"Eastlakes",
	"Eastwood",
	"Edensor Park",
	"Edgecliff",
	"Edmondson Park",
	"Elanora Heights",
	"Elderslie",
	"Elizabeth Bay",
	"Elizabeth Hills",
	"Ellis Lane",
	"Elvina Bay",
	"Emerton",
	"Emu Heights",
	"Emu Plains",
	"Enfield",
	"Engadine",
	"Englorie Park",
	"Enmore",
	"Epping",
	"Ermington",
	"Erskine Park",
	"Erskineville",
	"Eschol Park",
	"Eveleigh",
	"Fairfield",
	"Fairfield East",
	"Fairfield Heights",
	"Fairfield West",
	"Fairlight",
	"Fiddletown",
	"Five Dock",
	"Flemington",
	"Forest Glen",
	"Forest Lodge",
	"Forestville",
	"Freemans Reach",
	"Frenchs Forest",
	"Freshwater",
	"Galston",
	"Georges Hall",
	"Gilead",
	"Girraween",
	"Gladesville",
	"Glebe",
	"Gledswood Hills",
	"Glen Alpine",
	"Glendenning",
	"Glenfield",
	"Glenhaven",
	"Glenmore Park",
	"Glenorie",
	"Glenwood",
	"Glossodia",
	"Gordon",
	"Granville",
	"Grasmere",
	"Grays Point",
	"Great Mackerel Beach",
	"Green Valley",
	"Greenacre",
	"Greendale",
	"Greenfield Park",
	"Greenhills Beach",
	"Greenwich",
	"Gregory Hills",
	"Greystanes",
	"Guildford",
	"Guildford West",
	"Gymea",
	"Gymea Bay",
	"Haberfield",
	"Hammondville",
	"Harrington Park",
	"Harris Park",
	"Hassall Grove",
	"Hawkesbury River",
	"Haymarket",
	"Heathcote",
	"Hebersham",
	"Heckenberg",
	"Henley",
	"Hillsdale",
	"Hinchinbrook",
	"Hobartville",
	"Holroyd",
	"Holsworthy",
	"Homebush",
	"Homebush West",
	"Horningsea Park",
	"Hornsby",
	"Hornsby Heights",
	"Horsley Park",
	"Hoxton Park",
	"Hunters Hill",
	"Huntingwood",
	"Huntleys Cove",
	"Huntleys Point",
	"Hurlstone Park",
	"Hurstville",
	"Hurstville Grove",
	"Illawong",
	"Ingleburn",
	"Ingleside",
	"Jamisontown",
	"Jannali",
	"Jordan Springs",
	"Kangaroo Point",
	"Kareela",
	"Kearns",
	"Kellyville",
	"Kellyville Ridge",
	"Kemps Creek",
	"Kensington",
	"Kenthurst",
	"Kentlyn",
	"Killara",
	"Killarney Heights",
	"Kings Langley",
	"Kings Park",
	"Kingsford",
	"Kingsgrove",
	"Kingswood",
	"Kingswood Park",
	"Kirkham",
	"Kirrawee",
	"Kirribilli",
	"Kogarah",
	"Kogarah Bay",
	"Ku-ring-gai Chase",
	"Kurnell",
	"Kurraba Point",
	"Kyeemagh",
	"Kyle Bay",
	"La Perouse",
	"Lakemba",
	"Lalor Park",
	"Lane Cove",
	"Lane Cove North",
	"Lane Cove West",
	"Lansdowne",
	"Lansvale",
	"Laughtondale",
	"Lavender Bay",
	"Leets Vale",
	"Leichhardt",
	"Len Waters Estate",
	"Leonay",
	"Leppington",
	"Lethbridge Park",
	"Leumeah",
	"Lewisham",
	"Liberty Grove",
	"Lidcombe",
	"Lilli Pilli",
	"Lilyfield",
	"Lindfield",
	"Linley Point",
	"Little Bay",
	"Liverpool",
	"Llandilo",
	"Loftus",
	"Londonderry",
	"Long Point",
	"Longueville",
	"Lower Portland",
	"Luddenham",
	"Lugarno",
	"Lurnea",
	"Macquarie Fields",
	"Macquarie Links",
	"Macquarie Park",
	"Maianbar",
	"Malabar",
	"Manly",
	"Manly Vale",
	"Maraylya",
	"Marayong",
	"Maroota",
	"Maroubra",
	"Marrickville",
	"Marsden Park",
	"Marsfield",
	"Mascot",
	"Matraville",
	"Mays Hill",
	"McGraths Hill",
	"McMahons Point",
	"Meadowbank",
	"Melrose Park",
	"Menai",
	"Menangle Park",
	"Merrylands",
	"Merrylands West",
	"Middle Cove",
	"Middle Dural",
	"Middleton Grange",
	"Miller",
	"Millers Point",
	"Milperra",
	"Milsons Passage",
	"Milsons Point",
	"Minchinbury",
	"Minto",
	"Minto Heights",
	"Miranda",
	"Mona Vale",
	"Monterey",
	"Moore Park",
	"Moorebank",
	"Mortdale",
	"Mortlake",
	"Mosman",
	"Mount Annan",
	"Mount Colah",
	"Mount Druitt",
	"Mount Ku-ring-gai",
	"Mount Lewis",
	"Mount Pritchard",
	"Mount Vernon",
	"Mulgoa",
	"Mulgrave",
	"Narellan Vale",
	"Naremburn",
	"Narrabeen",
	"Narraweena",
	"Narwee",
	"Nelson",
	"Neutral Bay",
	"Newington",
	"Newport",
	"Newtown",
	"Normanhurst",
	"North Balgowlah",
	"North Bondi",
	"North Curl Curl",
	"North Epping",
	"North Manly",
	"North Narrabeen",
	"North Parramatta",
	"North Richmond",
	"North Rocks",
	"North Ryde",
	"North Seaforth",
	"North St Ives",
	"North St Marys",
	"North Strathfield",
	"North Sydney",
	"North Turramurra",
	"North Willoughby",
	"North Wahroonga",
	"Northbridge",
	"Northmead",
	"Northwood",
	"Oakhurst",
	"Oakville",
	"Oatlands",
	"Oatley",
	"Old Guildford",
	"Old Toongabbie",
	"Oran Park",
	"Orchard Hills",
	"Osborne Park",
	"Oxford Falls",
	"Oxley Park",
	"Oyster Bay",
	"Paddington",
	"Padstow",
	"Padstow Heights",
	"Pagewood",
	"Palm Beach",
	"Panania",
	"Parklea",
	"Parramatta",
	"Peakhurst",
	"Peakhurst Heights",
	"Pemulwuy",
	"Pendle Hill",
	"Pennant Hills",
	"Penrith",
	"Penshurst",
	"Petersham",
	"Phillip Bay",
	"Picnic Point",
	"Pitt Town",
	"Pleasure Point",
	"Plumpton",
	"Point Piper",
	"Port Botany",
	"Potts Hill",
	"Potts Point",
	"Prairiewood",
	"Prestons",
	"Prospect",
	"Punchbowl",
	"Putney",
	"Pymble",
	"Pyrmont",
	"Quakers Hill",
	"Queens Park",
	"Queenscliff",
	"Raby",
	"Ramsgate",
	"Ramsgate Beach",
	"Randwick",
	"Redfern",
	"Regents Park",
	"Regentville",
	"Revesby",
	"Revesby Heights",
	"Rhodes",
	"Richmond",
	"Riverstone",
	"Riverview",
	"Riverwood",
	"Rockdale",
	"The Rocks",
	"Rodd Point",
	"Rookwood",
	"Rooty Hill",
	"Ropes Crossing",
	"Rose Bay",
	"Rosebery",
	"Rosehill",
	"Roselands",
	"Rosemeadow",
	"Roseville",
	"Roseville Chase",
	"Rouse Hill",
	"Royal National Park",
	"Rozelle",
	"Ruse",
	"Rushcutters Bay",
	"Russell Lea",
	"Rydalmere",
	"Ryde",
	"Sackville North",
	"Sadleir",
	"Sandringham",
	"Sans Souci",
	"Scheyville",
	"Schofields",
	"Scotland Island",
	"Seaforth",
	"Sefton",
	"Seven Hills",
	"Shalvey",
	"Shanes Park",
	"Silverwater",
	"Singletons Mill",
	"Smeaton Grange",
	"Smithfield",
	"South Coogee",
	"South Hurstville",
	"South Maroota",
	"South Penrith",
	"South Turramurra",
	"South Wentworthville",
	"South Windsor",
	"Spring Farm",
	"St Andrews",
	"St Clair",
	"St Helens Park",
	"St Ives",
	"St Ives Chase",
	"St Johns Park",
	"St Leonards",
	"St Marys",
	"St Peters",
	"Stanhope Gardens",
	"Stanmore",
	"Strathfield",
	"Strathfield South",
	"Summer Hill",
	"Surry Hills",
	"Sutherland",
	"Sydenham",
	"Sydney Olympic Park",
	"Sylvania",
	"Sylvania Waters",
	"Tamarama",
	"Taren Point",
	"Telopea",
	"Tempe",
	"Tennyson Point",
	"Terrey Hills",
	"The Ponds",
	"The Rocks",
	"Thornleigh",
	"Toongabbie",
	"Tregear",
	"Turramurra",
	"Turrella",
	"Undercliffe",
	"Ultimo",
	"Varroville",
	"Vaucluse",
	"Villawood",
	"Vineyard",
	"Voyager Point",
	"Wahroonga",
	"Waitara",
	"Wakeley",
	"Wallacia",
	"Wareemba",
	"Warrawee",
	"Warriewood",
	"Warwick Farm",
	"Waterfall",
	"Waterloo",
	"Watsons Bay",
	"Wattle Grove",
	"Waverley",
	"Waverton",
	"Weavers",
	"Wedderburn",
	"Wentworth Point",
	"Wentworthville",
	"Werrington",
	"Werrington County",
	"Werrington Downs",
	"West Hoxton",
	"West Killara",
	"West Lindfield",
	"West Pennant Hills",
	"West Pymble",
	"West Ryde",
	"Westleigh",
	"Westmead",
	"Wetherill Park",
	"Whalan",
	"Wheeler Heights",
	"Wiley Park",
	"Willmot",
	"Willoughby",
	"Willoughby East",
	"Windsor",
	"Windsor Downs",
	"Winston Hills",
	"Wisemans Ferry",
	"Wolli Creek",
	"Wollstonecraft",
	"Woodbine",
	"Woodcroft",
	"Woodpark",
	"Woollahra",
	"Woolloomooloo",
	"Woolooware",
	"Woolwich",
	"Woronora",
	"Woronora Heights",
	"Yagoona",
	"Yarramundi",
	"Yarrawarrah",
	"Yennora",
	"Yowie Bay",
	"Zetland",
}
