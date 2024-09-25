import csv
import matplotlib.pyplot as plt

# Read data from CSV
stages = []
alloc = []
total_alloc = []
sys_mem = []
num_gc = []
heap_idle = []
freed = []

with open('redis_LRC_benchmark_memory_usage.csv', 'r') as file:
    reader = csv.DictReader(file)
    for row in reader:
        stages.append(row['Stage'])
        alloc.append(int(row['Allocated Memory (MiB)']))
        total_alloc.append(int(row['Total Allocated Memory (MiB)']))
        sys_mem.append(int(row['System Memory (MiB)']))
        num_gc.append(int(row['Number of Garbage Collections']))
        heap_idle.append(int(row['Idle Heap Memory (MiB)']))

# Plotting Allocated Memory
plt.figure(figsize=(10, 6))
plt.plot(stages, alloc, label='Allocated Memory (MiB)', marker='o')
plt.plot(stages, sys_mem, label='System Memory (MiB)', marker='o', linestyle='--')
plt.plot(stages, heap_idle, label='Idle Heap Memory (MiB)', marker='o', linestyle='dotted')

plt.xlabel('Stage')
plt.ylabel('Memory (MiB)')
plt.title('Memory Usage Before and After Garbage Collection')
plt.xticks(rotation=45, ha='right')
plt.legend()
plt.tight_layout()

# Save or display the plot
plt.savefig('memory_usage_plot.png')
plt.show()
