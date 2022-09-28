#### CRDTs
Conflict-free Replicated Data Types

This repo contains the implementation of some of the CRDTs as specified [here](https://hal.inria.fr/inria-00555588/document). This is an experimental repo. My goal here is to learn the properties of CRDTs, Vector Clocks, and other distributed systems concepts by implementing the primitives.

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
#### References

###### CRDTs
- [Conflict Free Replicated Data Type](https://en.wikipedia.org/wiki/Conflict-free_replicated_data_type)
- [CRDTs: example implementations in Go](https://github.com/neurodrone/crdt)
- [A Cmprehensive study of Convergent and Commutative Replicated Data Types](https://hal.inria.fr/inria-00555588/document)
- [A CRDT Primer Part I: Defanging Order Theory](http://jtfmumm.com/blog/2015/11/17/crdt-primer-1-defanging-order-theory/)
- [A CRDT Primer Part II: Convergent CRDTs](http://jtfmumm.com/blog/2015/11/24/crdt-primer-2-convergent-crdts/)
- [CRDTs go brrr](https://josephg.com/blog/crdts-go-brrr/)
- [Verifying Strong Eventual Consistency in Distributed
Systems](https://martin.kleppmann.com/papers/crdt-isabelle-oopsla17.pdf)
- [Roshi: A CRDT System For Timestamped Events by Sound Cloud](https://developers.soundcloud.com/blog/roshi-a-crdt-system-for-timestamped-events)
- [CRDTs: The Hard Parts](https://www.youtube.com/watch?v=x7drE24geUw)
- [CRDTs and the Quest for Distributed Consistency by Martin Kelppmann](https://www.youtube.com/watch?v=B5NULPSiOGw)


##### Vector Clocks (Logical time keeping)
- [Vector Clocks](https://sookocheff.com/post/time/vector-clocks/)
- [Data Laced With History](http://archagon.net/blog/2018/03/24/data-laced-with-history/)
- [Why Logical Clocks Are Easy](https://queue.acm.org/detail.cfm?id=2917756)
- [Keeping Time In Real Systems by Kavya Joshi](https://www.youtube.com/watch?v=BRvj8PykSc4)
- [How to Have your Causality and Wall Clocks, Too by Jon Moore](https://www.youtube.com/watch?v=YqNGbvFHoKM)
- [Living Without Atomic Clocks by Cockroach Labs](https://www.cockroachlabs.com/blog/living-without-atomic-clocks/)


