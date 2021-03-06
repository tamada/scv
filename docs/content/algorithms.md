---
title: ":gem: Algorithms"
---

## :speaking_head: Overview

`scv` supports simpson, jaccard, dice, cosine, pearson, euclidean, manhattan, and chebyshev algorithms.
Those algorithms separated into two categories, similarities, and distances.
The followings describes each algorithms and how to calculate similarity/distance from two vectors.

Before calculating the similarities and/or distance of the following algorithms,
we assume that two vectors were given (\\(v = \{ (a_1, b_1), (a_2, b_2), ..., (a_n, b_n) \}\\) and \\(w = \{ (c_1, d_1), (c_2, d_2), ..., (c_m, d_m) \}\\)).

## Similarities

Calculate similarities among the given vectors.
For this, `scv` computes similarity between two vectors from the given vectors of their combinations.

To describe each algorithm, 
Each element of the vector has key and value.

### Simpson Index

$$S = \frac{|\mathrm{intersect}(v, w)|}{\min(|v_1|, |v_2|)}$$

\\(\mathrm{intersect}\\) function returns the new vector by common keys and sum of thier values (\\(a_i = c_j (1 \leq i \leq n, 1 \leq j \leq m)\\).


### Jaccard index

$$J=\frac{|\mathrm{intersect}(v, w)|}{|\mathrm{union}(v, w)|}$$

\\(\mathrm{intersect}\\) function returns the new vector which contains every keys of \\(v\\) and \\(w\\).

### Dice index

$$D=\frac{2\times|\mathrm{intersect}(v, w)|}{|v| + |w|}$$

### Cosine similarity

$$C=\cos\theta=\frac{v\cdot w}{\sqrt{\sum_{i=0}^{n}b_i^2}\sqrt{\sum_{j=0}^{m}d_j^2}}$$

### Pearson correlation efficiency



## Distances

### Euclidean Distance



### Manhattan Distance



### Chebyshev Distance


### Edit Distance
