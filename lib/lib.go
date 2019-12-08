package lib

func BinarySearch(checkSlice []int, findVal int) int {
	pos := -1
	left, right := 0, len(checkSlice)  //此处right长度不减1 ， 如果最大值为查找值，此处减一代码进入死循环
	Loop:
		for {
			if(left >= right){
				break Loop
			}
			mid := (left + right) / 2
			switch true {
			case checkSlice[mid] < findVal :
				left = mid
			case checkSlice[mid] == findVal :
				pos = mid
				break Loop
			case checkSlice[mid] > findVal :
				right = mid
	
			}
	
		}

	return pos
  }

  func BinarySearch2(checkSlice []int,leftIndex, rightIndex, findVal int) int{
	mid := (leftIndex + rightIndex) / 2
	post := -1
	if rightIndex >= leftIndex {
		if checkSlice[mid] > findVal {
			post = BinarySearch2(checkSlice, leftIndex, mid - 1, findVal)
		}else if checkSlice[mid] < findVal {
			post = BinarySearch2(checkSlice, mid + 1, rightIndex, findVal)
		}else {
			post = mid
		}
	}
	return post
  }