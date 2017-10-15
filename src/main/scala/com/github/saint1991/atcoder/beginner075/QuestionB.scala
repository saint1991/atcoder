package com.github.saint1991.atcoder.beginner075

import com.github.saint1991.atcoder.util.{MineMap, ScannerApp}

object QuestionB extends ScannerApp {

  final val H = scanner.nextInt()
  final val W = scanner.nextInt()

  val mines = (0 until H).foldLeft(Seq.empty[(Int, Int)]) { (acc, h) =>
    acc ++  scanner.next().zipWithIndex.collect {
      case ('#', w) => (h, w)
    }
  }

  val mineMap = new MineMap(H, W)
  mineMap.setMines(mines:_*)

  val answer = mineMap.opened().map { line =>
    line.map(_.toString).mkString
  }.mkString("\n")

  println(answer)
}

