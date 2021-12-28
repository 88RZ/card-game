package main
import (
  "fmt"
  "sort"
)
func isHandOfStraights(hand []int, k int) bool{
		// Check if the number of cards in the hand is divisible by k. If not, we can’t create groups, so return false.
    if len(hand) % k != 0{
        return false
    }

    // Count the occurrences of each card in the given hand.
    count := make(map[int]int)
    for _, i := range hand{
        count[i] = count[i] + 1
    }
    
    // Sort the list and start traversing it from the lowest-ranking card. We can use a hash map by storing card numbers as keys and occurrences as values.
    sort.Ints(hand)
    i := 0
    n := len(hand)

    // Use a nested loop that runs k times.
    for i < n {
        current := hand[i]
        for j := 0; j < k; j++ {
        		// Check if the current card and the next k-1 cards (in increasing ranking) are in the count map. If any of them don’t exist, return false.
            if _, ok := count[current + j]; !ok || count[current + j] == 0 {
                return false
            }

            // When each of the required cards is found, decrease its number of occurrences in the count.
            count[current + j]--
        }

        // After a complete group is found, use a while loop to find the next group’s smallest card and determine which of the next cards in count has more than zero occurrences left.
        for i < n && count[hand[i]] == 0{
            i++
        }
    }

    // Return true if all cards are sorted into groups.
    return true
}

func main() {
    hand := []int{5,2,4,4,1,3,5,6,3}
    k := 3
    fmt.Println(isHandOfStraights(hand, k))
    
    hand2 := []int{1,9,3,5,7,4,2,9,11}
    k = 2
    fmt.Println(isHandOfStraights(hand2, k))
}