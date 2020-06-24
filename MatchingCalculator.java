
import java.lang.Math;
public class MatchingCalculator {

  int[][] costTable;
  public MatchingCalculator(int[][] costTab ){
    this.costTable=costTab;

  }

  static class Point{
    int x;
    int y;
    //boolean assigned;
    //boolean crossedout;

    Point(int a, int b){
      x=a;
      y=b;
      //assigned=false;
      //crossedout=false;
    }

  }

  public static MatchingCalculator calculateCost(Faces faceT){

    int n = faceT.frame1.length;
    //System.out.println("frame1 size is" + n );
    int[][] costTable = new int[n][n];
    Point[] letterpoints = new Point[n];
    Point[] numberpoints = new Point[n];

    int pos=0;
    while (pos<faceT.frame1.length){
      letterpoints[pos] = new Point(faceT.frame1[pos][0]+(faceT.frame1[pos][2]/2),faceT.frame1[pos][1]-faceT.frame1[pos][3]/2);
//      System.out.println("loop1  " + pos );
      pos++;
    }

    pos=0;
    while (pos<faceT.frame2.length){
      numberpoints[pos] = new Point(faceT.frame2[pos][0]+(faceT.frame2[pos][2]/2),faceT.frame2[pos][1]-faceT.frame2[pos][3]/2);
//      System.out.println("loop2  " + pos );
      pos++;
    }

    for(int i=0; i<letterpoints.length;i++){

      for(int j=0;j<numberpoints.length;j++){
        long result=(long) Math.sqrt(((letterpoints[i].x-numberpoints[j].x) * (letterpoints[i].x-numberpoints[j].x))  +   ((letterpoints[i].y-numberpoints[j].y) * (letterpoints[i].y-numberpoints[j].y)) );
        int res = (int)result;
        costTable[i][j]=res;
        System.out.println(res);
  //      System.out.println("loop3, i and j are  " + i + " " + j );
      }
    }

    MatchingCalculator euclid=new MatchingCalculator(costTable);
    return euclid;

  }

  }
