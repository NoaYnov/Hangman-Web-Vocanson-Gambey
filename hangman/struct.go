package hangman

type Hang struct {
	Loop     bool
	Win      bool
	Word     string
	Guess    []string
	José     []string
	NbLetter int
	NbTry    int
	I        int
	Letters  string
}

func (w *Hang) InitHang(diff string) {
	w.Loop = true
	w.Word = RandomWord(w.SaveFileToTab(diff))
	w.NbLetter = len(w.Word)
	for i := 0; i < len(w.Word); i++ {
		w.Guess = append(w.Guess, "_")
	}
	w.NbTry = 10
	w.José = SaveJoséToTab()
	w.I = 0
	w.Letters = ""
}
