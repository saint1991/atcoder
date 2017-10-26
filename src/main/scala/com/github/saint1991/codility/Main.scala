package com.github.saint1991.codility

object Main extends App {
  val a = 1000
  val binaryRepr = a.toBinaryString.split("1")

  val i = 0
  binaryRepr.foreach{ zeros =>
    println(zeros.length)
  }
}
