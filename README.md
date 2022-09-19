#### CRDTs
Conflict-free Replicated Data Types



### Run tests 
```go
go test ./...
```

### run benchmarks for vector clocks
```go
go test ./... -bench=.
```

#### References

###### CRDTs
- [Conflict Free Replicated Data Type](https://en.wikipedia.org/wiki/Conflict-free_replicated_data_type)
- http://jtfmumm.com/blog/2015/11/17/crdt-primer-1-defanging-order-theory/
- http://jtfmumm.com/blog/2015/11/24/crdt-primer-2-convergent-crdts/
- [CRDTs go brrr](https://josephg.com/blog/crdts-go-brrr/)
- https://martin.kleppmann.com/papers/crdt-isabelle-oopsla17.pdf
- https://hal.inria.fr/inria-00555588/document
- [Roshi: A CRDT System For Timestamped Events by Sound Cloud](https://developers.soundcloud.com/blog/roshi-a-crdt-system-for-timestamped-events)


##### Vector Clocks (Logical time keeping)
- [Vector Clocks](https://sookocheff.com/post/time/vector-clocks/)
- [Data Laced With History](http://archagon.net/blog/2018/03/24/data-laced-with-history/)
- [Why Logical Clocks Are Easy](https://queue.acm.org/detail.cfm?id=2917756)
- [Keeping Time In Real Systems by Kavya Joshi](https://www.youtube.com/watch?v=BRvj8PykSc4)
- [How to Have your Causality and Wall Clocks, Too by Jon Moore](https://www.youtube.com/watch?v=YqNGbvFHoKM)
- https://www.youtube.com/watch?v=x7drE24geUw
- https://www.youtube.com/watch?v=B5NULPSiOGw
- https://www.youtube.com/watch?v=4i7KrG1Afbk&list=UU_QIfHvN9auy2CoOdSfMWDw&index=596
- [Living Without Atomic Clocks by Cockroach Labs](https://www.cockroachlabs.com/blog/living-without-atomic-clocks/)


