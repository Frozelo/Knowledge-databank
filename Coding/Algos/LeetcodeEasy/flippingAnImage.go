// Given an n x n binary matrix image, flip the image horizontally, then invert it, and return the resulting image.
// To flip an image horizontally means that each row of the image is reversed.
// For example, flipping [1,1,0] horizontally results in [0,1,1].
// To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0.
// For example, inverting [0,1,1] results in [1,0,0].

// Example 1:

// Input: image = [[1,1,0],[1,0,1],[0,0,0]]
// Output: [[1,0,0],[0,1,0],[1,1,1]]
// Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
// Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]

// Example 2:

// Input: image = [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
// Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
// Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
// Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]

// O(n)
func flip(a int) int {
    if a == 0 {
        return 1
    }

    return 0
}

func flipAndInvertImage(image [][]int) [][]int {
    for row := 0; row < len(image); row++{
        left, right := 0, len(image[row])-1

        for left < right {
            image[row][left], image[row][right] = flip(image[row][right]), flip(image[row][left])
            left++
            right-- 
        }

        if left == right {
            image[row][right] = flip(image[row][right])
        }
    }

    return image
}
