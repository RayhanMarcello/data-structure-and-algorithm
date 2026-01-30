# selectionSort
# time (o(n^2))
# space (o(1))
c = [-5,3,2,1,-4,-4,7,2,2]

def selectionSort(arr):
    n =len(arr)
    for i in range (0, n):
        min_idx = i
        for j in range(i+1,n):
            if arr[j] < arr[min_idx]:
                min_idx = j
        arr[i], arr[min_idx] = arr[min_idx], arr[i]

selectionSort(c)
print(c)
