# insertonSort
# time (o(n^2))
# space (o(1))

b = [-5,3,2,1,-4,-4,7,2,2]

def inserton_sort(arr):
    n = len(arr)
    for i in range(1, n):
        for j in range(i, 0, -1):
            if arr[j-1] > arr[j]:
                arr[j-1], arr[j] = arr[j], arr[j-1]
            else:
                break

inserton_sort(b)
print(b)