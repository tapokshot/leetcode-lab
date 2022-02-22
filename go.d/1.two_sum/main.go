package main

func main() {

}

// brute force
func twoSumBruteForce(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[j] == target-nums[i] {
                return []int{i, j}
            }
        }
    }
    return []int{}
}

// Делаем через хешмапу
// 1. Получаем разницу между target и текущим значением nums
// 2. Проверяем есть ли такое число уже в хешпаме
// 3. Если есть то возвращаем его индекс первым, текущий индекс вторым
// 4. Если нету, то добавляем ключем в мапу текущее значение, а его индекс в значение ключа
func onePassHashTable(nums []int, target int) []int {
    var m = make(map[int]int)
    for i := 0; i < len(nums); i++ {
        compliment := target - nums[i]
        if _, ok := m[compliment]; ok {
            return []int{m[i], i}
        }
        m[nums[i]] = i
    }
    return []int{}
}

func twoPassHashTable(nums []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        m[nums[i]] = i
    }

    for i:= 0; i < len(nums); i++ {
        complement := target - nums[i]
        if v, ok := m[complement]; ok && v != i {
            return []int{i, m[complement]}
        }
    }
    return []int{}
}
