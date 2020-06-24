package main


import "time"
import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
  "strconv"
)

// func (b RowReducor) reduceRow(bidchannel chan <- int) {
//    feedback := make(chan AuctionItem)
//
//    s := make([]string, size)
//
//    for i := 0; i < size; i++

func (b RowReductor) Reduce(bidchannel chan <- int) {
  go func(b RowReductor) {
    fmt.Println("inside row reductor ", b.toreduce)
   //s := make([]string, b.size)
   min,_ := findMinAndMax(b.toreduce[b.RowNumber])

   for i :=0 ; i < (b.size); i++ {
     b.toreduce[b.RowNumber][i]=b.toreduce[b.RowNumber][i]-min
     }
     bidchannel <- 0
     fmt.Println("one reduction done  ", b.toreduce)
   }(b)

  // fmt.Println("reduction completed for one row  ", b.toreduce)
}


func (b ColReductor) Reduce(bidchannel chan <- int) {
  go func(b ColReductor) {
   s := make([]int, b.size)
   //min,_ = findMinAndMax(b.toreduce[i])

   for i :=0 ; i < (b.size); i++ {
     s[i]=b.toreduce[i][b.ColNumber]
     }

     fmt.Println("one column column is  ", s)
    min,_ := findMinAndMax(s)
    fmt.Println("minimum value is   ", min)

    for i :=0 ; i < (b.size); i++ {
      b.toreduce[i][b.ColNumber]=b.toreduce[i][b.ColNumber]-min
      }

     bidchannel <- 0
   }(b)
}



func step5( costTable [][]int, size int) [][]int{

  //int[][] costTable=inputArray;


  mark := make([][]bool, size)
  for i := 0; i < size; i++ {
        innerLen := i + 1
        mark[i] = make([]bool, size)
        for j := 0; j < innerLen; j++ {
            mark[i][j] = false
        }
    }

    assign := make([][]bool, size)
    for i := 0; i < size; i++ {
          innerLen := i + 1
          assign[i] = make([]bool, size)
          for j := 0; j < innerLen; j++ {
              assign[i][j] = false
          }
      }

      asterix := make([][]int, size)
      for i := 0; i < size; i++ {
            //innerLen := i + 1
            asterix[i] = make([]int, 2)
            for j := 0; j < 2; j++ {
                asterix[i][j] = 0
            }
        }

  //int[][] asterix=new int[costTable.length][2];
  c :=0;


  // System.out.println("executing step5 now  " );
  // for(int i=0;i<costTable.length;i++){
  //  System.out.println("");
  //   for(int j=0;j<costTable.length;j++){
  //     System.out.print(costTable[i][j]+  " ");
  //   }
  // }


  for a := 0; a < size; a++{
    for  b:= 0; b < size; b++{
      mark[a][b]=false
      assign[a][b]=false
    }
  }

  for i := 0; i < size; i++{
     first := true
     for j := 0; j < size; j++{
       if (costTable[i][j] == 0  &&  mark[i][j]==false ){
         if (first){
           //A(r,c) is assigned;
           first = false
           mark[i][j]=true
           assign[i][j]=true
           asterix[c][0]=i
           asterix[c][1]=j
           c++;


           for k := 0; k < size; k++{
             if (costTable[k][j]==0 && i!= k){
               mark[k][j]=true
               assign[k][j]=false
             }
           }
         }else{
           mark[i][j]=false
           assign[i][j]=false
         }
       }
     }
  }

  return asterix;
}



func step34(cost [][]int ,Size int){
  fmt.Println("Inside set 3 and 4")
  size := Size

  marked := make([][]bool, size)
  for i := 0; i < size; i++ {
        innerLen := i + 1
        marked[i] = make([]bool, size)
        for j := 0; j < innerLen; j++ {
            marked[i][j] = false
        }
    }

    assigned := make([][]bool, size)
    for i := 0; i < size; i++ {
          innerLen := i + 1
          assigned[i] = make([]bool, size)
          for j := 0; j < innerLen; j++ {
              assigned[i][j] = false
          }
      }



  tickr := make([]bool, size)
  for i := 0; i < size; i++ {
        tickr[i] = false
    }

    tickc := make([]bool, size)
    for i := 0; i < size; i++ {
          tickc[i] = false
      }

      linec := make([]bool, size)
      for i := 0; i < size; i++ {
            linec[i] = false
        }

        liner := make([]bool, size)
        for i := 0; i < size; i++ {
              liner[i] = false
          }
  newTick := true
  linecount := 0

  fmt.Println("tables marked assigned tick and line created ")

  for (linecount < size) {
    fmt.Println("inside   for (linecount < size loop ")
    // System.out.println("executing step 3 to 4 " );
    // for(int i=0;i<costTable.length;i++){
    //  System.out.println("");
    //   for(int j=0;j<costTable.length;j++){
    //     System.out.print(costTable[i][j]+  " ");
    //   }
    // }
  for i := 0; i < size; i++{
     first := true
     for j := 0; j < size; j++ {
       if (cost[i][j] == 0  &&  marked[i][j]==false ){
         if (first){
           //A(r,c) is assigned;
           first = false
           marked[i][j]=true
           assigned[i][j]=true
           for k := 0; k < size; k++{
             if (cost[k][j]==0 && i!= k){
               marked[k][j]=false
               assigned[k][j]=false
             }
           }
         }else{
           marked[i][j]=false
           assigned[i][j]=false
         }
       }
     }
  }

  for i := 0; i < size; i++{
    fmt.Println("tickr[i] set to true ")
    tickr[i] = true;
    for j := 0; j < size; j++{
      if (assigned[i][j]==true){
        tickr[i] = false
      }
    }

  }
  //}
 //System.out.println(" ");
 fmt.Println("P1: newtick set to false")
  newTick=false;



  // Ticking cols
  //boolean[] tick=new boolean[costTable.length];
  for (newTick){
    fmt.Println("inside for new tick")
  for i := 0; i < size; i++{
  if(!tickc[i]){
    for j := 0; j < size; j++{
      if (marked[i][j]==true){
        tickc[j] = true
         fmt.Println("P2: newtick set to true at location Pos 2: %v Pos 3: %v  " ,i,j);
        newTick=true
      }
    }

  }

  }
  // Ticking rows again
  //boolean[] tick=new boolean[costTable.length];
  for i := 0; i < size; i++{
  if(!tickr[i]){
    for j := 0; j < size; j++{
      if (assigned [i][j]==true){
        tickr[j] = true
           fmt.Println("P3: newtick set to true at location Pos 2: %v Pos 3: %v ",i,j);
        newTick=true
      }
    }

  }

  }
  }

  minline := 0;
  for i := 0; i < size; i++{
  if(!tickr[i]){
    liner[i]=true;
    minline=minline+1

  }else{
    liner[i]=false
  }
  }

  for i := 0; i < size; i++{
  if(!tickc[i]){
    linec[i]=true
    minline=minline+1
  }else{
    linec[i]=false
  }
  }
  linecount=minline



  ///step 4 here
  smallest :=0
  firstfound :=false


  //fmt.Println("executing step 4  only " );
  // for(int i=0;i<costTable.length;i++){
  //  System.out.println("");
  //   for(int j=0;j<costTable.length;j++){
  //     System.out.print(costTable[i][j]+  " ");
  //   }
  // }
 fmt.Println("befor=e we start  step 4", cost)
 fmt.Println("linecount before step4 is ", linecount)
  if(linecount<size){

    fmt.Println("finding smallest int in step 4")
    for i := 0; i < size; i++{
      for j := 0; j < size; j++{
        if(!firstfound && assigned[i][j]==false){
          smallest=cost[i][j];
          firstfound=true
        }else{
          if(firstfound && cost[i][j]<smallest && assigned[i][j]==false ){
            smallest=cost[i][j]
          }
        }

      }
    }
//
  fmt.Println("substracting smallest from unmarked")
    for i := 0; i < size; i++{
      for j := 0; j < size; j++{
        if(marked[i][j]==false){
          cost[i][j]=cost[i][j]-smallest
        }
      }
    }

fmt.Println("adding smallest to intersections ")
    for i := 0; i < size; i++{
      for j := 0; j < size; j++{
        if(liner[i] && linec[j]){
          cost[i][j]=cost[i][j]+smallest
        }

      }
    }

    // for (int i=0;i<costTable.length; i++){
    //   liner[i]=false;
    //   linec[j]=false;
    //   tickr[i]=false;
    //   tickc[j]=false;
    //   for (int i=0;i<costTable.length; i++){
    //     marked[i][j]==false;
    //     assigned[i][j]==false
    //   }
    // }
  }



}


}

//}





func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

type ColReductor struct{
  toreduce [][]int
  ColNumber int
  size int

}
func NewColReductor( tabToreduce [][]int,col int ,Size int ) ColReductor {
  //tabtoreduce=tabToreduce;
  p := ColReductor{toreduce: tabToreduce,ColNumber: col, size: Size}
  return p
}

type RowReductor struct{
  toreduce [][]int
  RowNumber int
  size int
}
func NewRowReductor( tabToreduce [][]int,line int ,Size int ) RowReductor {
  p := RowReductor{toreduce: tabToreduce,RowNumber: line, size: Size}
  return p
}

func (a Reducable) isAllReduced() bool {
  return a.state.reduced
}
func (a Reducable) setAllReduced(value bool) {
   a.state.reduced = value
}

func NewReducable() Reducable {
  s:= State{ reduced:false}
	f := Reducable{allReduced: false, state: &s}
  return f
}

type Reducable struct {
  allReduced bool
  state *State
}

type State struct {
  reduced bool
}





func ReduceManager(size int, receivechannel chan int,reducable *Reducable) {
   go func() {
    allRowsReduced := false
    var count int =0
    //fmt.Println("inside the reduce manager  ")
    //var completed bool =false
    for (!allRowsReduced){
    //  fmt.Println("Auction house waiting for message.")
      notification := <- receivechannel
      if(notification==0){
        count=count+1
        //fmt.Println("increasing count for reduced rows ")
        if(count == size){
          allRowsReduced=true

          reducable.setAllReduced(true)
          //fmt.Println("allreduced was set to true ", reducable.isAllReduced())
        }

      }

    }
   }()
 }

func main() {
	// Open the file
	csvfile, err := os.Open("cost_matrix_3.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records


  	// Read each record from csv
    count :=1
    size :=0
for {
  	record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
    if (count > 1){
    //  fmt.Printf(a,b,record[3])
      fmt.Printf("Pos 1: %v Pos 2: %v Pos 3: %v \n",  record[1],record[2],record[3])
      size=size+1


    }
    count=count+1
	}
  //fmt.Printf("Size is %v", size)

  //on cree le tableau contena nt les valeur dela costmatrix
  //var costs [size][size]int
  costs := make([][]int, size)
  for i := 0; i < size; i++ {
        innerLen := i + 1
        costs[i] = make([]int, size)
        for j := 0; j < innerLen; j++ {
            costs[i][j] = 0
        }
    }

    posletter := make([]string, size)
    for i := 0; i < size; i++ {
          posletter[i]=" "
      }

      posnumber := make([]string, size)
      for i := 0; i < size; i++ {
            posnumber[i]=" "
        }

  csvfile, err = os.Open("cost_matrix_3.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
  r = csv.NewReader(csvfile)
  lineNumber :=0
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records

	for {
    //fmt.Println("inside my for loop ")
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
      //fmt.Println("about to break")
			break
		}
    //fmt.Println("did not  break")
		if err != nil {
			log.Fatal(err)
		}

    if(lineNumber == 0){

      for j := 0; j < size; j++ {
        var res string =record[lineNumber][j]
            posletter[j] = res
        }
    }

    fmt.Println("letter positions ", posletter)

    //fmt.Println("about to start getting elements ")
    if(lineNumber>0){
      //fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
      for j := 0; j < size; j++ {
        res, _ := strconv.Atoi(record[j+1])
        //fmt.Printf("l-1: %v j: %v \n", (lineNumber-1), j)
            costs[lineNumber-1][j] = res
            //fmt.Println("element at pos ", res)
        }
    }
    lineNumber = lineNumber+1
	}
  fmt.Println("costs ", costs)
//on a finit de remplir le tableau initial de cost


//   fmt.Println("before row reductiuon ", costs)
// for i := 0; i < size; i++ {
//   res, _ := strconv.Atoi(record[i+1])
//   //fmt.Printf("l-1: %v j: %v \n", (lineNumber-1), j)
//       costs[lineNumber-1][i] = res
//       //fmt.Println("element at pos ", res)
//   }
  rowReducers := make([]RowReductor, size)
  for i := 0; i < size; i++ {
        rowReducers[i]=NewRowReductor(costs, i, size)
    }

    //fmt.Println("after we create rowreducers ", costs)

    notificationChannel := make(chan int)
    reducable := NewReducable()
    ReduceManager(size, notificationChannel, &reducable)
    for i := 0; i < size; i++ {
          rowReducers[i].Reduce(notificationChannel)
      }





    for !reducable.isAllReduced(){
        time.Sleep(10 * time.Millisecond)
        //fmt.Println("isAllreduced still false  ")
      }
    //  fmt.Println("after  reducing rows ", costs)

  fmt.Println("Row reduction is done and we have  ", costs)
    //about to rudeuce colums
    colReducers := make([]ColReductor, size)
    for i := 0; i < size; i++ {
          colReducers[i]=NewColReductor(costs, i, size)
      }
    notificationChannel2 := make(chan int)
    reducable2 := NewReducable()
    ReduceManager(size, notificationChannel2, &reducable2)
    for i := 0; i < size; i++ {
          colReducers[i].Reduce(notificationChannel2)
      }

      for !reducable2.isAllReduced(){
          time.Sleep(1 * time.Millisecond)
        }

    fmt.Println("Column reduction is done and we have  ", costs)

    step34(costs, size)

    fmt.Println("After step 3 and 4  ", costs)

    res :=step5(costs, size)

    fmt.Println("res is   ", res)







}
