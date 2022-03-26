package scylla_test

import (
	"context"
	"testing"
	"time"

	"github.com/Posrabi/flashy/backend/apitest"
	"github.com/Posrabi/flashy/backend/users/pkg/entity"
	"github.com/Posrabi/flashy/backend/users/pkg/repository"
	"github.com/Posrabi/flashy/backend/users/pkg/scylla"
	"github.com/stretchr/testify/require"
)

func TestPhraseRepository(t *testing.T) {
	phraseSetup(t)

	t.Run("Create_Phrase", func(t *testing.T) {
		testCreate_Phrase(t, apitest.PhraseRepo)
	})

	t.Run("Get_Phrases", func(t *testing.T) {
		testGet_Phrases(t, apitest.PhraseRepo)
	})

	t.Run("Delete_Phrase", func(t *testing.T) {
		testDelete_Phrase(t, apitest.PhraseRepo)
	})
}

func phraseSetup(t *testing.T) {
	sess := apitest.Setup(t)

	apitest.PhraseRepo = scylla.NewPhraseRepository(sess)
	t.Cleanup(func() {
		for _, phrase := range apitest.TestPhrases {
			require.NoError(t, apitest.PhraseRepo.DeletePhrase(context.Background(), phrase.UserID, phrase.Time))
		}
	})
}

func testCreate_Phrase(t *testing.T, repo repository.Phrase) {
	t.Helper()

	for _, phrase := range apitest.TestPhrases {
		require.NoError(t, repo.CreatePhrase(context.Background(), phrase))
	}
}

func testGet_Phrases(t *testing.T, repo repository.Phrase) {
	t.Helper()

	for _, phrase := range apitest.TestPhrases {
		phrases, err := repo.GetPhrases(context.Background(), phrase.UserID, time.Now().Add(time.Minute*-1), time.Now())
		require.NoError(t, err)
		require.Equal(t, phrase.Word, phrases[0].Word)
		require.Equal(t, phrase.Sentence, phrases[0].Sentence)
	}
}

func testDelete_Phrase(t *testing.T, repo repository.Phrase) {
	t.Helper()

	for i, phrase := range apitest.TestPhrases {
		require.NoError(t, repo.DeletePhrase(context.Background(), phrase.UserID, phrase.Time))
		removePhraseAtIndex(apitest.TestPhrases, i)
	}
}

func removePhraseAtIndex(phrases []*entity.Phrase, i int) []*entity.Phrase {
	return append(phrases[:i], phrases[i+1:]...)
}
