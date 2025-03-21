package savannah

import (
	"math/rand"
	"time"
)

func GetMagic8BallResponse(round int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return Magic8Ball[r.Intn(len(Magic8Ball))]
}

var Magic8Ball = []string{
	ItIsCertain,
	YesDefinitely,
	ReplyHazyTryAgain,
	AskAgainLater,
	DontCountOnIt,
}

var (
	ItIsCertain       = "It is certain."
	ReplyHazyTryAgain = "Reply hazy, try again."
	DontCountOnIt     = "Don't count on it."
	YesDefinitely     = "Yes – definitely."
	AskAgainLater     = "Ask again later."
)
