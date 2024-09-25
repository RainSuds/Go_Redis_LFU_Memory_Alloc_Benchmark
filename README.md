# Go Redis LFU Memory Allocation Benchmark
This is a mini benchmark that is similar to the redis LFU cache measuring the memory allocation of each operations.

# Run
To compile the go code after you clone the repository
```
$go run .
```
Testing with the first test file
```
Input a file: test/test1.txt
```
Run python script
```
$python3 plot_data.py
```
# Output
![alt text](https://github.com/RainSuds/Go_Redis_LFU_Memory_Alloc_Benchmark/blob/main/memory_usage_plot.png?raw=true)
