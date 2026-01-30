# buuble sort 
# time(o(n^2)), 
# space(n(1))

a = [-5,3,2,1,-3,-3,7,2,2]

print("before bubble sort", a)

def bubble_sort(arr):
    n = len(arr)
    swap = True
    while swap:
        swap = False
        for i in range(1, n):
            if arr[i-1] > arr[i]:
                swap = True
                arr[i-1], arr[i] = arr[i], arr[i-1]

bubble_sort(a)
print("after bubble sort", a)
