package com.github.saint1991.atcoder.util

// Representations of height x width Minesweeper map to solve opened state when the place of mines are given
class MineMap(val height: Int, val width: Int) {

  type Mine = (Int, Int)

  private val internalCounts = Array.fill(height)(Array.fill(width)(0))
  private var mines = Seq.empty[Mine]

  private def isInside(coordinate: Mine): Boolean =
    0 <= coordinate._1 && coordinate._1 < height && 0 <= coordinate._2 && coordinate._2 < width

  def setMines(mines: Mine*): Unit = mines.foreach { mine =>
    this.mines = this.mines :+ mine
    for {
      h <- mine._1 - 1 to mine._1 + 1
      w <- mine._2 - 1 to mine._2 + 1
    } yield {
      if (isInside((h, w))) internalCounts(h)(w) += 1
    }
  }

  def opened(mineChar: String = "#"): Array[Array[String]] = {
    val stringMap = internalCounts.map { line =>
      line.map(_.toString)
    }
    this.mines.foreach { mine =>
      stringMap(mine._1)(mine._2) = mineChar
    }
    stringMap
  }
}
