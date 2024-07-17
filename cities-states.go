package n_american_cities

type city struct {
	name, provState, country string
}

type stateProv struct {
	name, abbrev string
}

var cities = []city{
	// Canada
	{"Airdrie", "AB", "CAN"},
	{"Calgary", "AB", "CAN"},
	{"Edmonton", "AB", "CAN"},
	{"Grande Prairie", "AB", "CAN"},
	{"Lethbridge", "AB", "CAN"},
	{"Medicine Hat", "AB", "CAN"},
	{"Red Deer", "AB", "CAN"},
	{"St. Albert", "AB", "CAN"},
	{"Strathcona County", "AB", "CAN"},
	{"Wood Buffalo", "AB", "CAN"},
	{"Abbotsford", "BC", "CAN"},
	{"Burnaby", "BC", "CAN"},
	{"Chilliwack", "BC", "CAN"},
	{"Coquitlam", "BC", "CAN"},
	{"Delta", "BC", "CAN"},
	{"Kamloops", "BC", "CAN"},
	{"Kelowna", "BC", "CAN"},
	{"Langley", "BC", "CAN"},
	{"Maple Ridge", "BC", "CAN"},
	{"Nanaimo", "BC", "CAN"},
	{"New Westminster", "BC", "CAN"},
	{"North Vancouver", "BC", "CAN"},
	{"Port Coquitlam", "BC", "CAN"},
	{"Prince George", "BC", "CAN"},
	{"Richmond", "BC", "CAN"},
	{"Saanich", "BC", "CAN"},
	{"Surrey", "BC", "CAN"},
	{"Vancouver", "BC", "CAN"},
	{"Victoria", "BC", "CAN"},
	{"Winnipeg", "MB", "CAN"},
	{"Fredericton", "NB", "CAN"},
	{"Moncton", "NB", "CAN"},
	{"Saint John", "NB", "CAN"},
	{"St. John's", "NL", "CAN"},
	{"Cape Breton", "NS", "CAN"},
	{"Halifax", "NS", "CAN"},
	{"Ajax", "ON", "CAN"},
	{"Aurora", "ON", "CAN"},
	{"Barrie", "ON", "CAN"},
	{"Belleville", "ON", "CAN"},
	{"Brampton", "ON", "CAN"},
	{"Brantford", "ON", "CAN"},
	{"Burlington", "ON", "CAN"},
	{"Caledon", "ON", "CAN"},
	{"Cambridge", "ON", "CAN"},
	{"Chatham-Kent", "ON", "CAN"},
	{"Clarington", "ON", "CAN"},
	{"Greater Sudbury", "ON", "CAN"},
	{"Guelph", "ON", "CAN"},
	{"Halton Hills", "ON", "CAN"},
	{"Hamilton", "ON", "CAN"},
	{"Kawartha Lakes", "ON", "CAN"},
	{"Kingston", "ON", "CAN"},
	{"Kitchener", "ON", "CAN"},
	{"London", "ON", "CAN"},
	{"Markham", "ON", "CAN"},
	{"Milton", "ON", "CAN"},
	{"Mississauga", "ON", "CAN"},
	{"Newmarket", "ON", "CAN"},
	{"Niagara Falls", "ON", "CAN"},
	{"Norfolk County", "ON", "CAN"},
	{"North Bay", "ON", "CAN"},
	{"Oakville", "ON", "CAN"},
	{"Oshawa", "ON", "CAN"},
	{"Ottawa", "ON", "CAN"},
	{"Peterborough", "ON", "CAN"},
	{"Pickering", "ON", "CAN"},
	{"Richmond Hill", "ON", "CAN"},
	{"Sarnia", "ON", "CAN"},
	{"Sault Ste. Marie", "ON", "CAN"},
	{"St. Catharines", "ON", "CAN"},
	{"Thunder Bay", "ON", "CAN"},
	{"Toronto", "ON", "CAN"},
	{"Vaughan", "ON", "CAN"},
	{"Waterloo", "ON", "CAN"},
	{"Welland", "ON", "CAN"},
	{"Whitby", "ON", "CAN"},
	{"Windsor", "ON", "CAN"},
	{"Blainville", "QC", "CAN"},
	{"Brossard", "QC", "CAN"},
	{"Drummondville", "QC", "CAN"},
	{"Gatineau", "QC", "CAN"},
	{"Granby", "QC", "CAN"},
	{"Laval", "QC", "CAN"},
	{"Longueuil", "QC", "CAN"},
	{"Lévis", "QC", "CAN"},
	{"Mirabel", "QC", "CAN"},
	{"Montreal", "QC", "CAN"},
	{"Quebec City", "QC", "CAN"},
	{"Repentigny", "QC", "CAN"},
	{"Saguenay", "QC", "CAN"},
	{"Saint-Hyacinthe", "QC", "CAN"},
	{"Saint-Jean-sur-Richelieu", "QC", "CAN"},
	{"Saint-Jérôme", "QC", "CAN"},
	{"Sherbrooke", "QC", "CAN"},
	{"Terrebonne", "QC", "CAN"},
	{"Trois-Rivières", "QC", "CAN"},
	{"Regina", "SK", "CAN"},
	{"Saskatoon", "SK", "CAN"},
	//Mexico
	{"Mexico City", "Mexico City", "MEX"},
	{"Ecatepec", "State of Mexico", "MEX"},
	{"León", "Guanajuato", "MEX"},
	{"Puebla", "Puebla", "MEX"},
	{"Ciudad Juárez", "Chihuahua", "MEX"},
	{"Guadalajara", "Jalisco", "MEX"},
	{"Zapopan", "Jalisco", "MEX"},
	{"Monterrey", "Nuevo León", "MEX"},
	{"Nezahualcóyotl", "State of Mexico", "MEX"},
	{"Chihuahua", "Chihuahua", "MEX"},
	{"Mérida", "Yucatán", "MEX"},
	{"Naucalpan", "State of Mexico", "MEX"},
	{"Cancún", "Quintana Roo", "MEX"},
	{"Saltillo", "Coahuila", "MEX"},
	{"Aguascalientes", "Aguascalientes", "MEX"},
	{"Hermosillo", "Sonora", "MEX"},
	{"Mexicali", "Baja California", "MEX"},
	{"San Luis Potosí", "San Luis Potosí", "MEX"},
	{"Culiacán", "Sinaloa", "MEX"},
	{"Querétaro", "Querétaro", "MEX"},
	{"Morelia", "Michoacán", "MEX"},
	{"Chimalhuacán", "State of Mexico", "MEX"},
	{"Reynosa", "Tamaulipas", "MEX"},
	{"Torreón", "Coahuila", "MEX"},
	{"Tlalnepantla", "State of Mexico", "MEX"},
	{"Acapulco", "Guerrero", "MEX"},
	{"Tlaquepaque", "Jalisco", "MEX"},
	{"Guadalupe", "Nuevo León", "MEX"},
	{"Durango", "Durango", "MEX"},
	{"Tuxtla Gutiérrez", "Chiapas", "MEX"},
	{"Cuautitlán Izcalli", "State of Mexico", "MEX"},
	{"Veracruz", "Veracruz", "MEX"},
	{"Ciudad Apodaca", "Nuevo León", "MEX"},
	{"Ciudad López Mateos", "State of Mexico", "MEX"},
	{"Matamoros", "Tamaulipas", "MEX"},
	{"General Escobedo", "Nuevo León", "MEX"},
	{"Irapuato", "Guanajuato", "MEX"},
	{"Xalapa", "Veracruz", "MEX"},
	{"Tonalá", "Jalisco", "MEX"},
	{"Mazatlán", "Sinaloa", "MEX"},
	{"Nuevo Laredo", "Tamaulipas", "MEX"},
	{"San Nicolás de los Garza", "Nuevo León", "MEX"},
	{"Ojo de Agua", "State of Mexico", "MEX"},
	{"Xico", "State of Mexico", "MEX"},
	{"Celaya", "Guanajuato", "MEX"},
	{"Tepic", "Nayarit", "MEX"},
	{"Ixtapaluca", "State of Mexico", "MEX"},
	{"Cuernavaca", "Morelos", "MEX"},
	{"Villahermosa", "Tabasco", "MEX"},
	{"Ciudad Victoria", "Tamaulipas", "MEX"},
	{"Ensenada", "Baja California", "MEX"},
	{"Ciudad Obregón", "Sonora", "MEX"},
	{"Ciudad Nicolás Romero", "State of Mexico", "MEX"},
	{"Soledad", "San Luis Potosí", "MEX"},
	{"Ciudad Benito Juárez", "Nuevo León", "MEX"},
	{"Playa del Carmen", "Quintana Roo", "MEX"},
	{"Santa Catarina", "Nuevo León", "MEX"},
	{"Gómez Palacio", "Durango", "MEX"},
	{"Uruapan", "Michoacán", "MEX"},
	{"Los Mochis", "Sinaloa", "MEX"},
	{"Pachuca", "Hidalgo", "MEX"},
	{"Tampico", "Tamaulipas", "MEX"},
	{"Tehuacán", "Puebla", "MEX"},
	{"San Francisco Coacalco", "State of Mexico", "MEX"},
	{"Nogales", "Sonora", "MEX"},
	{"Oaxaca", "Oaxaca", "MEX"},
	{"La Paz", "Baja California Sur", "MEX"},
	{"Campeche", "Campeche", "MEX"},
	{"Monclova", "Coahuila", "MEX"},
	{"García", "Nuevo León", "MEX"},
	{"Chilpancingo", "Guerrero", "MEX"},
	{"Puerto Vallarta", "Jalisco", "MEX"},
	{"Toluca", "State of Mexico", "MEX"},
	{"Tapachula", "Chiapas", "MEX"},
	{"Buenavista", "State of Mexico", "MEX"},
	{"Coatzacoalcos", "Veracruz", "MEX"},
	{"Ciudad Madero", "Tamaulipas", "MEX"},
	{"Cabo San Lucas", "Baja California Sur", "MEX"},
	{"Chicoloapan", "State of Mexico", "MEX"},
	{"Ciudad del Carmen", "Campeche", "MEX"},
	{"San Cristóbal de las Casas", "Chiapas", "MEX"},
	{"Poza Rica", "Veracruz", "MEX"},
	{"San Juan del Río", "Querétaro", "MEX"},
	{"San Luis Río Colorado", "Sonora", "MEX"},
	{"Chalco", "State of Mexico", "MEX"},
	{"Jiutepec", "Morelos", "MEX"},
	{"Piedras Negras", "Coahuila", "MEX"},
	{"Guadalupe", "Zacatecas", "MEX"},
	{"Chetumal", "Quintana Roo", "MEX"},
	{"Miramar", "Tamaulipas", "MEX"},
	{"Salamanca", "Guanajuato", "MEX"},
	{"Ciudad Acuña", "Coahuila", "MEX"},
	{"Manzanillo", "Colima", "MEX"},
	{"San Pablo de las Salinas", "State of Mexico", "MEX"},
	{"Cuautla", "Morelos", "MEX"},
	{"Zamora", "Michoacán", "MEX"},
	{"Minatitlán", "Veracruz", "MEX"},
	{"Villa de Álvarez", "Colima", "MEX"},
	{"Colima", "Colima", "MEX"},
	// USA
	{"Anchorage", "AK", "USA"},
	{"Birmingham", "AL", "USA"},
	{"Huntsville", "AL", "USA"},
	{"Mobile", "AL", "USA"},
	{"Montgomery", "AL", "USA"},
	{"Tuscaloosa", "AL", "USA"},
	{"Fayetteville", "AR", "USA"},
	{"Little Rock", "AR", "USA"},
	{"Buckeye", "AZ", "USA"},
	{"Chandler", "AZ", "USA"},
	{"Gilbert", "AZ", "USA"},
	{"Glendale", "AZ", "USA"},
	{"Goodyear", "AZ", "USA"},
	{"Mesa", "AZ", "USA"},
	{"Peoria", "AZ", "USA"},
	{"Phoenix", "AZ", "USA"},
	{"Scottsdale", "AZ", "USA"},
	{"Surprise", "AZ", "USA"},
	{"Tempe", "AZ", "USA"},
	{"Tucson", "AZ", "USA"},
	{"Yuma", "AZ", "USA"},
	{"Anaheim", "CA", "USA"},
	{"Antioch", "CA", "USA"},
	{"Bakersfield", "CA", "USA"},
	{"Berkeley", "CA", "USA"},
	{"Burbank", "CA", "USA"},
	{"Carlsbad", "CA", "USA"},
	{"Chico", "CA", "USA"},
	{"Chula Vista", "CA", "USA"},
	{"Clovis", "CA", "USA"},
	{"Concord", "CA", "USA"},
	{"Corona", "CA", "USA"},
	{"Costa Mesa", "CA", "USA"},
	{"Downey", "CA", "USA"},
	{"El Cajon", "CA", "USA"},
	{"El Monte", "CA", "USA"},
	{"Elk Grove", "CA", "USA"},
	{"Escondido", "CA", "USA"},
	{"Fairfield", "CA", "USA"},
	{"Fontana", "CA", "USA"},
	{"Fremont", "CA", "USA"},
	{"Fresno", "CA", "USA"},
	{"Fullerton", "CA", "USA"},
	{"Garden Grove", "CA", "USA"},
	{"Glendale", "CA", "USA"},
	{"Hayward", "CA", "USA"},
	{"Hesperia", "CA", "USA"},
	{"Huntington Beach", "CA", "USA"},
	{"Inglewood", "CA", "USA"},
	{"Irvine", "CA", "USA"},
	{"Jurupa Valley", "CA", "USA"},
	{"Lancaster", "CA", "USA"},
	{"Long Beach", "CA", "USA"},
	{"Los Angeles", "CA", "USA"},
	{"Menifee", "CA", "USA"},
	{"Modesto", "CA", "USA"},
	{"Moreno Valley", "CA", "USA"},
	{"Murrieta", "CA", "USA"},
	{"Oakland", "CA", "USA"},
	{"Oceanside", "CA", "USA"},
	{"Ontario", "CA", "USA"},
	{"Orange", "CA", "USA"},
	{"Oxnard", "CA", "USA"},
	{"Palmdale", "CA", "USA"},
	{"Pasadena", "CA", "USA"},
	{"Pomona", "CA", "USA"},
	{"Rancho Cucamonga", "CA", "USA"},
	{"Rialto", "CA", "USA"},
	{"Richmond", "CA", "USA"},
	{"Riverside", "CA", "USA"},
	{"Roseville", "CA", "USA"},
	{"Sacramento", "CA", "USA"},
	{"Salinas", "CA", "USA"},
	{"San Bernardino", "CA", "USA"},
	{"San Diego", "CA", "USA"},
	{"San Francisco", "CA", "USA"},
	{"San Jose", "CA", "USA"},
	{"San Mateo", "CA", "USA"},
	{"Santa Ana", "CA", "USA"},
	{"Santa Clara", "CA", "USA"},
	{"Santa Clarita", "CA", "USA"},
	{"Santa Maria", "CA", "USA"},
	{"Santa Rosa", "CA", "USA"},
	{"Simi Valley", "CA", "USA"},
	{"Stockton", "CA", "USA"},
	{"Sunnyvale", "CA", "USA"},
	{"Temecula", "CA", "USA"},
	{"Thousand Oaks", "CA", "USA"},
	{"Torrance", "CA", "USA"},
	{"Vacaville", "CA", "USA"},
	{"Vallejo", "CA", "USA"},
	{"Ventura", "CA", "USA"},
	{"Victorville", "CA", "USA"},
	{"Visalia", "CA", "USA"},
	{"West Covina", "CA", "USA"},
	{"Arvada", "CO", "USA"},
	{"Aurora", "CO", "USA"},
	{"Boulder", "CO", "USA"},
	{"Centennial", "CO", "USA"},
	{"Colorado Springs", "CO", "USA"},
	{"Denver", "CO", "USA"},
	{"Fort Collins", "CO", "USA"},
	{"Greeley", "CO", "USA"},
	{"Lakewood", "CO", "USA"},
	{"Pueblo", "CO", "USA"},
	{"Thornton", "CO", "USA"},
	{"Westminster", "CO", "USA"},
	{"Bridgeport", "CT", "USA"},
	{"Hartford", "CT", "USA"},
	{"New Haven", "CT", "USA"},
	{"Stamford", "CT", "USA"},
	{"Waterbury", "CT", "USA"},
	{"Washington", "DC", "USA"},
	{"Cape Coral", "FL", "USA"},
	{"Clearwater", "FL", "USA"},
	{"Coral Springs", "FL", "USA"},
	{"Davie", "FL", "USA"},
	{"Fort Lauderdale", "FL", "USA"},
	{"Gainesville", "FL", "USA"},
	{"Hialeah", "FL", "USA"},
	{"Hollywood", "FL", "USA"},
	{"Jacksonville", "FL", "USA"},
	{"Lakeland", "FL", "USA"},
	{"Miami", "FL", "USA"},
	{"Miami Gardens", "FL", "USA"},
	{"Miramar", "FL", "USA"},
	{"Orlando", "FL", "USA"},
	{"Palm Bay", "FL", "USA"},
	{"Palm Coast", "FL", "USA"},
	{"Pembroke Pines", "FL", "USA"},
	{"Pompano Beach", "FL", "USA"},
	{"Port St. Lucie", "FL", "USA"},
	{"St. Petersburg", "FL", "USA"},
	{"Tallahassee", "FL", "USA"},
	{"Tampa", "FL", "USA"},
	{"West Palm Beach", "FL", "USA"},
	{"Athens", "GA", "USA"},
	{"Atlanta", "GA", "USA"},
	{"Augusta", "GA", "USA"},
	{"Columbus", "GA", "USA"},
	{"Macon", "GA", "USA"},
	{"Sandy Springs", "GA", "USA"},
	{"Savannah", "GA", "USA"},
	{"South Fulton", "GA", "USA"},
	{"Honolulu", "HI", "USA"},
	{"Cedar Rapids", "IA", "USA"},
	{"Davenport", "IA", "USA"},
	{"Des Moines", "IA", "USA"},
	{"Boise", "ID", "USA"},
	{"Meridian", "ID", "USA"},
	{"Nampa", "ID", "USA"},
	{"Aurora", "IL", "USA"},
	{"Chicago", "IL", "USA"},
	{"Elgin", "IL", "USA"},
	{"Joliet", "IL", "USA"},
	{"Naperville", "IL", "USA"},
	{"Peoria", "IL", "USA"},
	{"Rockford", "IL", "USA"},
	{"Springfield", "IL", "USA"},
	{"Carmel", "IN", "USA"},
	{"Evansville", "IN", "USA"},
	{"Fishers", "IN", "USA"},
	{"Fort Wayne", "IN", "USA"},
	{"Indianapolis", "IN", "USA"},
	{"South Bend", "IN", "USA"},
	{"Kansas City", "KS", "USA"},
	{"Olathe", "KS", "USA"},
	{"Overland Park", "KS", "USA"},
	{"Topeka", "KS", "USA"},
	{"Wichita", "KS", "USA"},
	{"Lexington", "KY", "USA"},
	{"Louisville", "KY", "USA"},
	{"Baton Rouge", "LA", "USA"},
	{"Lafayette", "LA", "USA"},
	{"New Orleans", "LA", "USA"},
	{"Shreveport", "LA", "USA"},
	{"Boston", "MA", "USA"},
	{"Brockton", "MA", "USA"},
	{"Cambridge", "MA", "USA"},
	{"Lowell", "MA", "USA"},
	{"Lynn", "MA", "USA"},
	{"New Bedford", "MA", "USA"},
	{"Quincy", "MA", "USA"},
	{"Springfield", "MA", "USA"},
	{"Worcester", "MA", "USA"},
	{"Baltimore", "MD", "USA"},
	{"Ann Arbor", "MI", "USA"},
	{"Dearborn", "MI", "USA"},
	{"Detroit", "MI", "USA"},
	{"Grand Rapids", "MI", "USA"},
	{"Lansing", "MI", "USA"},
	{"Sterling Heights", "MI", "USA"},
	{"Warren", "MI", "USA"},
	{"Minneapolis", "MN", "USA"},
	{"Rochester", "MN", "USA"},
	{"Saint Paul", "MN", "USA"},
	{"Columbia", "MO", "USA"},
	{"Independence", "MO", "USA"},
	{"Kansas City", "MO", "USA"},
	{"Lee's Summit", "MO", "USA"},
	{"Springfield", "MO", "USA"},
	{"St. Louis", "MO", "USA"},
	{"Jackson", "MS", "USA"},
	{"Billings", "MT", "USA"},
	{"Cary", "NC", "USA"},
	{"Charlotte", "NC", "USA"},
	{"Concord", "NC", "USA"},
	{"Durham", "NC", "USA"},
	{"Fayetteville", "NC", "USA"},
	{"Greensboro", "NC", "USA"},
	{"High Point", "NC", "USA"},
	{"Raleigh", "NC", "USA"},
	{"Wilmington", "NC", "USA"},
	{"Winston-Salem", "NC", "USA"},
	{"Fargo", "ND", "USA"},
	{"Lincoln", "NE", "USA"},
	{"Omaha", "NE", "USA"},
	{"Manchester", "NH", "USA"},
	{"Edison", "NJ", "USA"},
	{"Elizabeth", "NJ", "USA"},
	{"Jersey City", "NJ", "USA"},
	{"Lakewood", "NJ", "USA"},
	{"Newark", "NJ", "USA"},
	{"Paterson", "NJ", "USA"},
	{"Woodbridge", "NJ", "USA"},
	{"Albuquerque", "NM", "USA"},
	{"Las Cruces", "NM", "USA"},
	{"Rio Rancho", "NM", "USA"},
	{"Henderson", "NV", "USA"},
	{"Las Vegas", "NV", "USA"},
	{"North Las Vegas", "NV", "USA"},
	{"Reno", "NV", "USA"},
	{"Sparks", "NV", "USA"},
	{"Albany", "NY", "USA"},
	{"Buffalo", "NY", "USA"},
	{"New York", "NY", "USA"},
	{"Rochester", "NY", "USA"},
	{"Syracuse", "NY", "USA"},
	{"Yonkers", "NY", "USA"},
	{"Akron", "OH", "USA"},
	{"Cincinnati", "OH", "USA"},
	{"Cleveland", "OH", "USA"},
	{"Columbus", "OH", "USA"},
	{"Dayton", "OH", "USA"},
	{"Toledo", "OH", "USA"},
	{"Broken Arrow", "OK", "USA"},
	{"Norman", "OK", "USA"},
	{"Oklahoma City", "OK", "USA"},
	{"Tulsa", "OK", "USA"},
	{"Bend", "OR", "USA"},
	{"Eugene", "OR", "USA"},
	{"Gresham", "OR", "USA"},
	{"Hillsboro", "OR", "USA"},
	{"Portland", "OR", "USA"},
	{"Salem", "OR", "USA"},
	{"Allentown", "PA", "USA"},
	{"Philadelphia", "PA", "USA"},
	{"Pittsburgh", "PA", "USA"},
	{"Providence", "RI", "USA"},
	{"Charleston", "SC", "USA"},
	{"Columbia", "SC", "USA"},
	{"North Charleston", "SC", "USA"},
	{"Sioux Falls", "SD", "USA"},
	{"Chattanooga", "TN", "USA"},
	{"Clarksville", "TN", "USA"},
	{"Knoxville", "TN", "USA"},
	{"Memphis", "TN", "USA"},
	{"Murfreesboro", "TN", "USA"},
	{"Nashville", "TN", "USA"},
	{"Abilene", "TX", "USA"},
	{"Allen", "TX", "USA"},
	{"Amarillo", "TX", "USA"},
	{"Arlington", "TX", "USA"},
	{"Austin", "TX", "USA"},
	{"Beaumont", "TX", "USA"},
	{"Brownsville", "TX", "USA"},
	{"Carrollton", "TX", "USA"},
	{"College Station", "TX", "USA"},
	{"Conroe", "TX", "USA"},
	{"Corpus Christi", "TX", "USA"},
	{"Dallas", "TX", "USA"},
	{"Denton", "TX", "USA"},
	{"Edinburg", "TX", "USA"},
	{"El Paso", "TX", "USA"},
	{"Fort Worth", "TX", "USA"},
	{"Frisco", "TX", "USA"},
	{"Garland", "TX", "USA"},
	{"Grand Prairie", "TX", "USA"},
	{"Houston", "TX", "USA"},
	{"Irving", "TX", "USA"},
	{"Killeen", "TX", "USA"},
	{"Laredo", "TX", "USA"},
	{"League City", "TX", "USA"},
	{"Lewisville", "TX", "USA"},
	{"Lubbock", "TX", "USA"},
	{"McAllen", "TX", "USA"},
	{"McKinney", "TX", "USA"},
	{"Mesquite", "TX", "USA"},
	{"Midland", "TX", "USA"},
	{"New Braunfels", "TX", "USA"},
	{"Odessa", "TX", "USA"},
	{"Pasadena", "TX", "USA"},
	{"Pearland", "TX", "USA"},
	{"Plano", "TX", "USA"},
	{"Richardson", "TX", "USA"},
	{"Round Rock", "TX", "USA"},
	{"San Antonio", "TX", "USA"},
	{"Sugar Land", "TX", "USA"},
	{"Tyler", "TX", "USA"},
	{"Waco", "TX", "USA"},
	{"Wichita Falls", "TX", "USA"},
	{"Provo", "UT", "USA"},
	{"Salt Lake City", "UT", "USA"},
	{"St. George", "UT", "USA"},
	{"West Jordan", "UT", "USA"},
	{"West Valley City", "UT", "USA"},
	{"Alexandria", "VA", "USA"},
	{"Chesapeake", "VA", "USA"},
	{"Hampton", "VA", "USA"},
	{"Newport News", "VA", "USA"},
	{"Norfolk", "VA", "USA"},
	{"Richmond", "VA", "USA"},
	{"Suffolk", "VA", "USA"},
	{"Virginia Beach", "VA", "USA"},
	{"Bellevue", "WA", "USA"},
	{"Everett", "WA", "USA"},
	{"Kent", "WA", "USA"},
	{"Renton", "WA", "USA"},
	{"Seattle", "WA", "USA"},
	{"Spokane", "WA", "USA"},
	{"Spokane Valley", "WA", "USA"},
	{"Tacoma", "WA", "USA"},
	{"Vancouver", "WA", "USA"},
	{"Green Bay", "WI", "USA"},
	{"Madison", "WI", "USA"},
	{"Milwaukee", "WI", "USA"},
}

var (
	canadianProvs = []stateProv{
		{"Ontario", "ON"},
		{"Quebec", "QC"},
		{"Nova Scotia", "NS"},
		{"New Brunswick", "NB"},
		{"Manitoba", "MB"},
		{"British Columbia", "BC"},
		{"Prince Edward Island", "PE"},
		{"Saskatchewan", "SK"},
		{"Alberta", "AB"},
		{"Newfoundland and Labrador", "NL"},
	}

	// Mexican states appear to not have abbreviations

	usStates = []stateProv{
		{"Alabama", "AL"},
		{"Alaska", "AK"},
		{"Arizona", "AZ"},
		{"Arkansas", "AR"},
		{"California", "CA"},
		{"Colorado", "CO"},
		{"Connecticut", "CT"},
		{"Delaware", "DE"},
		{"Florida", "FL"},
		{"Georgia", "GA"},
		{"Hawaii", "HI"},
		{"Idaho", "ID"},
		{"Illinois", "IL"},
		{"Indiana", "IN"},
		{"Iowa", "IA"},
		{"Kansas", "KS"},
		{"Kentucky", "KY"},
		{"Louisiana", "LA"},
		{"Maine", "ME"},
		{"Maryland", "MD"},
		{"Massachusetts", "MA"},
		{"Michigan", "MI"},
		{"Minnesota", "MN"},
		{"Mississippi", "MS"},
		{"Missouri", "MO"},
		{"Montana", "MT"},
		{"Nebraska", "NE"},
		{"Nevada", "NV"},
		{"New Hampshire", "NH"},
		{"New Jersey", "NJ"},
		{"New Mexico", "NM"},
		{"New York", "NY"},
		{"North Carolina", "NC"},
		{"North Dakota", "ND"},
		{"Ohio", "OH"},
		{"Oklahoma", "OK"},
		{"Oregon", "OR"},
		{"Pennsylvania", "PA"},
		{"Rhode Island", "RI"},
		{"South Carolina", "SC"},
		{"South Dakota", "SD"},
		{"Tennessee", "TN"},
		{"Texas", "TX"},
		{"Utah", "UT"},
		{"Vermont", "VT"},
		{"Virginia", "VA"},
		{"Washington", "WA"},
		{"West Virginia", "WV"},
		{"Wisconsin", "WI"},
		{"Wyoming", "WY"},

		{"District of Columbia", "DC"},
	}
)
