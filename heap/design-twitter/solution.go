package designtwitter

import "container/heap"

type tweet struct {
	tweetID   int
	previous  *tweet
	timestamp int
}

type Twitter struct {
	count   int
	tweets  map[int]*tweet
	follows map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{
		count:   0,
		tweets:  make(map[int]*tweet),
		follows: make(map[int]map[int]bool),
	}
}

func (t *Twitter) PostTweet(userID int, tweetID int) {
	t.userInit(userID)

	t.count++
	t.tweets[userID] = &tweet{
		tweetID:   tweetID,
		previous:  t.tweets[userID],
		timestamp: t.count,
	}
}

func (t *Twitter) GetNewsFeed(userID int) []int {
	t.userInit(userID)

	feed, h := make([]int, 0), &tweetHeap{}
	heap.Init(h)

	for folleweeID, isFollowing := range t.follows[userID] {
		if isFollowing {
			t.userInit(folleweeID)
			heap.Push(h, t.tweets[folleweeID])
		}
	}

	for h.Len() > 0 && len(feed) < 10 {
		next, _ := heap.Pop(h).(*tweet)

		if next.timestamp == 0 {
			break
		}

		feed = append(feed, next.tweetID)
		heap.Push(h, next.previous)
	}

	return feed
}

func (t *Twitter) Follow(followerID int, followeeID int) {
	t.userInit(followerID)

	t.follows[followerID][followeeID] = true
}

func (t *Twitter) Unfollow(followerID int, followeeID int) {
	t.userInit(followerID)

	t.follows[followerID][followeeID] = false
}

func (t *Twitter) userInit(id int) {
	if _, found := t.tweets[id]; found {
		return
	}

	t.tweets[id] = &tweet{}
	t.follows[id] = make(map[int]bool)
	t.follows[id][id] = true
}

type tweetHeap []*tweet

func (tweetHeap tweetHeap) Len() int {
	return len(tweetHeap)
}

func (tweetHeap tweetHeap) Less(i, j int) bool {
	return tweetHeap[i].timestamp > tweetHeap[j].timestamp
}

func (tweetHeap tweetHeap) Swap(i, j int) {
	tweetHeap[i], tweetHeap[j] = tweetHeap[j], tweetHeap[i]
}

func (tweetHeap *tweetHeap) Push(x any) {
	val, _ := x.(*tweet)

	*tweetHeap = append(*tweetHeap, val)
}

func (tweetHeap *tweetHeap) Pop() any {
	current := *tweetHeap
	value := current[len(current)-1]
	*tweetHeap = current[:len(current)-1]

	return value
}
