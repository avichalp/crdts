#### CRDTs
Conflict-free Replicated Data Types

This repo contains a few implementations of CRDTs and logical clocks. It also has various resources on the problem of clock synchronization in distributed systems. My goal with this is to learn the properties of CRDTs, Vector Clocks, and other distributed systems concepts by implementing the primitives.




--- 

Current status

- [G counter](https://github.com/avichalp/crdts/blob/master/crdts/gcounter.go)
- [PN counter](https://github.com/avichalp/crdts/blob/master/crdts/pncounter.go)
- [G set](https://github.com/avichalp/crdts/blob/master/crdts/gset.go)
- [2 Phase set](https://github.com/avichalp/crdts/blob/master/crdts/2pset.go)
- [LWW set](https://github.com/avichalp/crdts/blob/master/crdts/lwwset.go)
- [OR set](https://github.com/avichalp/crdts/blob/master/crdts/orset.go)
- [Vector Clocks](https://github.com/avichalp/crdts/tree/master/vector_clocks)
- [Vector Clocks based 2 Phase set](https://github.com/avichalp/crdts/blob/master/crdts/v2pset.go)

---

### Run tests 
```go
go test ./...
```

### run benchmarks for vector clocks
```go
go test ./... -bench=.
```

---
### References

#### CRDTs, Vector Clocks & Distributed Timekeeping

##### Articles, Blog posts
- [Conflict Free Replicated Data Type](https://en.wikipedia.org/wiki/Conflict-free_replicated_data_type)
- [A Cmprehensive study of Convergent and Commutative Replicated Data Types](https://hal.inria.fr/inria-00555588/document)
- [A CRDT Primer Part I: Defanging Order Theory](http://jtfmumm.com/blog/2015/11/17/crdt-primer-1-defanging-order-theory/)
- [A CRDT Primer Part II: Convergent CRDTs](http://jtfmumm.com/blog/2015/11/24/crdt-primer-2-convergent-crdts/)
- [CRDTs go brrr](https://josephg.com/blog/crdts-go-brrr/)
- [Verifying Strong Eventual Consistency in Distributed
Systems](https://martin.kleppmann.com/papers/crdt-isabelle-oopsla17.pdf)
- [Vector Clocks](https://sookocheff.com/post/time/vector-clocks/)
- [Data Laced With History](http://archagon.net/blog/2018/03/24/data-laced-with-history/)
- [Why Logical Clocks Are Easy](https://queue.acm.org/detail.cfm?id=2917756)
- [Living Without Atomic Clocks by Cockroach Labs](https://www.cockroachlabs.com/blog/living-without-atomic-clocks/)
- [Spanner Paper](https://storage.googleapis.com/pub-tools-public-publication-data/pdf/65b514eda12d025585183a641b5a9e096a3c4be5.pdf)
- [Spanner, TrueTime and CAP](https://research.google/pubs/pub45855/)
- [Cockroach Labs's Consistency Models](https://www.cockroachlabs.com/blog/consistency-model/)
- [Dynamo Paper](https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf)
- [Practical uses of synchronized clocks in distributed systems - Barbara Liskov](https://dl.acm.org/doi/pdf/10.1145/112600.112601)
- [Notes collected from various papers on CRDTs and Eventual Consistency](https://github.com/pfrazee/crdt_notes)

##### Videos
- [CRDTs: The Hard Parts](https://www.youtube.com/watch?v=x7drE24geUw)
- [CRDTs and the Quest for Distributed Consistency by Martin Kelppmann](https://www.youtube.com/watch?v=B5NULPSiOGw)
- [Keeping Time In Real Systems by Kavya Joshi](https://www.youtube.com/watch?v=BRvj8PykSc4)
- [How to Have your Causality and Wall Clocks, Too by Jon Moore](https://www.youtube.com/watch?v=YqNGbvFHoKM)


##### Reference implementations
- [CRDTs: example implementations in Go](https://github.com/neurodrone/crdt)
- [Roshi: A CRDT System For Timestamped Events by Sound Cloud](https://developers.soundcloud.com/blog/roshi-a-crdt-system-for-timestamped-events)
- [Vector Clocks implementation in Go](https://github.com/aprimadi/vector-clock)

