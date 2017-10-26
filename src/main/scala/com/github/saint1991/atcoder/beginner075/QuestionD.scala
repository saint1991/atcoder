package com.github.saint1991.atcoder.beginner075

import com.github.saint1991.atcoder.util.ScannerApp

case class Coordinate(x: Long, y: Long) {
  def distanceTo(another: Coordinate): Long = Math.sqrt(x * x + y * y).toLong
}

object QuestionD extends ScannerApp {

  final val N = scanner.nextInt()
  final val K = scanner.nextInt()

  val coordinates: Seq[Coordinate] = for {
    _ <- 0 until N
    x = scanner.nextLong()
    y = scanner.nextLong()
  } yield Coordinate(x, y)

  val distanceMap: Seq[((Int, Int), Long)] = (for {
    i <- 0 until N
    j <- 0 until N
    c1 = coordinates(i)
    c2 = coordinates(j)
  } yield (i, j) -> c1.distanceTo(c2)).sortBy(_._2)

  val pair = distanceMap.par.find { case ((i, j), _) =>
    if (i == j) false
    else {
      val c1 = coordinates(i)
      val c2 = coordinates(j)

      val left = Math.min(c1.x, c2.x)
      val right = Math.max(c1.x, c2.x)
      val top = Math.max(c1.y, c2.y)
      val bottom = Math.min(c1.y, c2.y)

      val k = coordinates.count{ c => left <= c.x && c.x <= right && bottom <= c.y && c.y <= top }
      k >= K
    }
  }.get._1

  // 残念，必ずしも頂点がcoordinatesの2点とは限らん
  val c1 = coordinates(pair._1)
  val c2 = coordinates(pair._2)
  val answer: Long = Math.abs(c1.x - c2.x) * Math.abs(c1.y - c2.y)
  println(answer)

}
