package hangman

type Hang struct {
	Loop     bool
	Word     string
	Guess    []string
	José     []string
	NbLetter int
	NbTry    int
	I        int
	letters  string
}

func (w *Hang) InitHang(dif string) {
	w.Loop = true
	w.Word = RandomWord(w.SaveFileToTab(dif))
	w.NbLetter = len(w.Word)
	for i := 0; i < len(w.Word); i++ {
		w.Guess = append(w.Guess, "_")
 	}
 	w.NbTry = 10
 	w.José = SaveJoséToTab()
}
