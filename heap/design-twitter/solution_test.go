package designtwitter_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	designtwitter "github.com/iktzdx/go-neetcode/heap/design-twitter"
)

func Test_Twitter(t *testing.T) {
	tw := designtwitter.Constructor()
	require.NotZero(t, tw, "create a new instance of Twitter")

	// User 1 posts a new tweet (id = 5).
	tw.PostTweet(1, 5)

	// User 1's news feed should return a list with 1 tweet id -> [5].
	f1 := tw.GetNewsFeed(1)
	require.True(t, reflect.DeepEqual([]int{5}, f1), "get user 1's news feed")

	// User 2 posts a new tweet (id = 6).
	tw.PostTweet(2, 6)

	// User 2's news feed should return a list with 1 tweet id -> [6].
	f2 := tw.GetNewsFeed(2)
	require.True(t, reflect.DeepEqual([]int{6}, f2), "get user 2's news feed")

	// User 1 follows user 2.
	tw.Follow(1, 2)

	// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
	// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	f1 = tw.GetNewsFeed(1)
	require.True(t, reflect.DeepEqual([]int{6, 5}, f1), "get user 1's news feed after follow")

	// User 2's news feed should still only contain their own tweets -> [6].
	f2 = tw.GetNewsFeed(2)
	require.True(t, reflect.DeepEqual([]int{6}, f2), "get user 2's news feed after follow")

	// User 1 unfollows user 2.
	tw.Unfollow(1, 2)

	// User 1's news feed should return a list with 1 tweet id -> [5],
	// since user 1 is no longer following user 2.
	f1 = tw.GetNewsFeed(1)
	require.True(t, reflect.DeepEqual([]int{5}, f1), "get user 1's news feed after unfollow")

	tweetsID := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

	// User 3 posts a new tweet.
	for _, id := range tweetsID {
		tw.PostTweet(3, id)
	}

	// Should return tweets ID in descending order.
	sort.Slice(tweetsID, func(i, j int) bool {
		return tweetsID[i] > tweetsID[j]
	})

	// User 1 follows user 3.
	tw.Follow(1, 3)

	// User 1's news feed should return a list with 10 recent tweet IDs.
	f1 = tw.GetNewsFeed(1)
	require.True(t, reflect.DeepEqual(tweetsID[:10], f1), "get 10 recent tweet IDs from user 1's news feed")
}
