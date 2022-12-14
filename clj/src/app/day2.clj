(ns app.day2 (:require [clojure.string :as str]
                       [clojure.java.io :as io]
                       [criterium.core :as crit]))

(def opts1
  (let [loss 0 draw 3 win 6 rock 1 paper 2 scissor 3]
    {0 0
     "A X" (+ rock draw)
     "A Y" (+ paper win)
     "A Z" (+ scissor loss)

     "B X" (+ rock loss)
     "B Y" (+ paper draw)
     "B Z" (+ scissor win)

     "C X" (+ rock win)
     "C Y" (+ paper loss)
     "C Z" (+ scissor draw)}))

(def opts2 
  (let [loss 0 draw 3 win 6 rock 1 paper 2 scissor 3]
    {0 0
     "A Y" (+ rock draw)
     "A Z" (+ paper win)
     "A X" (+ scissor loss)

     "B X" (+ rock loss)
     "B Y" (+ paper draw)
     "B Z" (+ scissor win)

     "C Z" (+ rock win)
     "C X" (+ paper loss)
     "C Y" (+ scissor draw)}))

(defn alg [rules lines]
  (let [line-scores (map rules lines)]
    (reduce + line-scores)))

(defn alg2 [rules lines]
    (reduce (fn [a b] (+ a (rules b))) 0 lines))

(comment

  (let [input (slurp (io/resource "day_2"))
        lines (str/split-lines input)]
    {:part1 (alg opts1 lines)
     :part2 (alg opts2 lines)})
 
  (let [input (slurp (io/resource "day_2"))
        lines (str/split-lines input)]
    (crit/quick-bench (alg opts1 lines))
    (crit/quick-bench (alg opts2 lines)))

  (let [input (slurp (io/resource "day_2"))
        lines (str/split-lines input)]
    (crit/quick-bench (alg2 opts1 lines)))



  )