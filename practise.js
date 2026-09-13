console.time("test")

for(let i  = 0 ; i <= 1000000; i++){
    console.log(i)
}

const end = performance.now()

console.timeEnd("test")