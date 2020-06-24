import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class FaceTracking {
  // int[][] frame1;
  // int[][] frame2;
  //
  // public FaceTracking(int[][] f1, int[][] f2){
  //   this.frame1=f1;
  //   this.frame2=f2;
  // }

  public static void main(String[] args) throws FileNotFoundException {
        //Get scanner instance
        //System.out.println("args[0]: " + args[0]);
        //System.out.println("args[1]: " + args[1]);
        Scanner scanner = new Scanner(new File(args[0]));
        //Set the delimiter used in file
        scanner.useDelimiter(",");

        //Get all tokens and store them in some data structure
        //I am just printing them
        int rowcount=0;
        int items=0;
        while (scanner.hasNext())
        {
            System.out.print(scanner.next() + "|");
            items++;
            if(items==4){
              rowcount++;
              items=0;
            }
        }
      //  System.out.print( "Row count :"+ rowcount);
         int[][] frame1 = new int[rowcount][4];
         int[][] frame2 = new int[rowcount][4];


        scanner = new Scanner(new File(args[0]));
        scanner.useDelimiter(",");
        int i=0;
        int j=0;

//while loop to fill out first array of frame
        while (scanner.hasNext())
        {
          String temp = scanner.next();

        //  System.out.println("P1: "+temp);
            if(j==0){
              j++;
            }
            else{
              frame1[i][j-1]=Integer.parseInt(temp);
              j++;
              if(j==4){
                i++;
                rowcount++;
                j=0;
              }
            }


        }

        i=0;
        j=0;
        scanner = new Scanner(new File(args[1]));
        scanner.useDelimiter(",");
        while (scanner.hasNext())
        {

          String temp = scanner.next();

          //System.out.println("P2: "+temp);
            if(j==0){
              j++;



            }
            else{
              frame2[i][j-1]=Integer.parseInt(temp);
              j++;
              if(j==4){
                i++;
                rowcount++;
                j=0;
              }
            }
        }

        scanner.close();
      //  System.out.println("done reading the files  " );
        Faces faceTracker = new Faces(frame1,frame2);
    //    System.out.println("about to calculate cost " );
        MatchingCalculator euclid= MatchingCalculator.calculateCost(faceTracker);
        MatchingCalculator step1out= Hungarian.reduceRow(euclid);
        MatchingCalculator step2out= Hungarian.reduceColumn(step1out);
        int[][] step5out= Hungarian.optimal(step2out);

        for (int c=0;c<step5out.length; c++){
          System.out.println("result number "+ 1 + "is "+ step5out[c][0] + " " +  step5out[c][1] );

        }



        //Do not forget to close the scanner

    }

}
