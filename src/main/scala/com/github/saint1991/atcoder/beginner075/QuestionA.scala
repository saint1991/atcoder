package com.github.saint1991.atcoder.beginner075

import com.github.saint1991.atcoder.util.ScannerApp

object QuestionA extends ScannerApp {

  val a = scanner.nextInt()
  val b = scanner.nextInt()
  val c = scanner.nextInt()

  val answer = Seq(a, b, c).foldLeft(Map[Int, Int]()) { (acc, element) =>
    acc.updated(element, acc.getOrElse(element, 0) + 1)
  }.minBy(_._2)._1

  println(answer)
}
