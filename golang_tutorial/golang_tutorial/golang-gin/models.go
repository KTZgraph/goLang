package main

// Structura karty

type Card struct {
	ID    int    `json:"id" binding:"required"`
	Type  int    `json:"typ"`                      //0 karta pytanie, 1 karta odpowiedź
	Blank int    `json:"puste" binding:"required"` //ile kart odpowiedzi na pytanie
	Text  string `json:"tekst" binding:"required"` // podłoga to luka
}

// Lista card
var cards = []Card{
	Card{0, 0, 0, "Adoptowanie dziecka tylko po, by porzucić je w galerii handlowej."},
	Card{1, 0, 0, "Sieroty, które wybuchają gdy ktoś je tylko pokocha."},
	Card{2, 0, 0, "Niekończąca się aktualizacja Windowsa"},
	Card{3, 0, 0, "Bug."},
	Card{4, 0, 0, "Wadliwe testy jednostkowe."},
	Card{5, 0, 0, "Użyszkodnik."},
	Card{6, 0, 0, "Aplikacja kalkulatora zajmująca 5GB."},
	Card{7, 1, 1, "Nie programuję w święta, bo _ się rodzi."},
	Card{8, 1, 1, "3 rzeczy lepsze od seksu. _ _ _"},
	Card{9, 0, 0, "Audyt bezpiezeństwa teleinformatycznego."},
	Card{10, 0, 0, "Testy na produkcji."},
	Card{11, 0, 0, "Pasywno-agresywna notatka o bałaganie w mieszkaniu."},
	Card{12, 0, 0, "Spanko."},
	Card{13, 0, 0, "Granko."},
	Card{14, 0, 0, "Niech żyje zbrodniczy reżim."},
	Card{15, 0, 0, "Kuce z Orła."},
	Card{16, 0, 0, "Wałówka od mamy."},
	Card{17, 0, 0, "10 dni do wypłaty."},
	Card{18, 0, 0, "Bo to zła kobieta była."},
	Card{19, 1, 2, "Rude to _ i _."},
	Card{20, 0, 0, "Tłusty drewnojad."},
	Card{21, 0, 0, "Zupa z paczki o smaku opakowania."},
	Card{22, 0, 0, "Przez twe oczy zielone, zielone _."},
	Card{23, 0, 0, "Naleśnikowa środa."},
	Card{24, 0, 0, "Kończenie dokumentacji o 5 nad ranem."},
	Card{25, 0, 0, "Powtarzanie roku."},
	Card{26, 0, 0, "Warunek."},
	Card{27, 0, 0, "Pralka."},
	Card{28, 0, 0, "Chodzenie po starym rynku w szpilkach."},
	Card{29, 0, 0, "Nieudana migracja bez logów."},
	Card{30, 0, 0, "Nocny dyżur."},
	Card{31, 1, 1, "Kac morderca jest wtedy gdy _."},
	Card{32, 0, 0, "Buka."},
	Card{33, 0, 0, "Nie mogę wyjść, mam _."},
	Card{34, 0, 0, "Cukier."},
	Card{35, 0, 0, "Co?"},
	Card{36, 0, 0, "Chodź, debata jest."},
	Card{37, 0, 0, "Ranga w pubg."},
	Card{38, 0, 0, "Dla Ciebie Pani Sarna."},
	Card{39, 1, 1, "Palenie papierosów powoduje _."},
	Card{40, 0, 0, "Syndrom Sztokholmski."},
	Card{41, 0, 0, "Domyśl się."},
}
