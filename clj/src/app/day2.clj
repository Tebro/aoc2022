(ns app.day2 (:require [clojure.string :as str]
                       [clojure.java.io :as io]))

(defn opts [] 
  (let [loss 0 draw 3 win 6 rock 1 paper 2 scissor 3]
    {"A X" (+ rock draw)
     "A Y" (+ paper win)
     "A Z" (+ scissor loss)

     "B X" (+ rock loss)
     "B Y" (+ paper draw)
     "B Z" (+ scissor win)

     "C X" (+ rock win)
     "C Y" (+ paper loss)
     "C Z" (+ scissor draw)}))

(defn opts2 [] 
  (let [loss 0 draw 3 win 6 rock 1 paper 2 scissor 3]
    {"A Y" (+ rock draw)
     "A Z" (+ paper win)
     "A X" (+ scissor loss)

     "B X" (+ rock loss)
     "B Y" (+ paper draw)
     "B Z" (+ scissor win)

     "C Z" (+ rock win)
     "C X" (+ paper loss)
     "C Y" (+ scissor draw)}))

(defn day2 []
  (let [scores (opts)
        scores2 (opts2)
        input (slurp (io/resource "day_2"))
        lines (str/split-lines input)
        line-scores (map #(hash-map :score1 (scores %) :score2 (scores2 %)) lines)]
    {:first (reduce + (map :score1 line-scores))
     :second (reduce + (map :score2 line-scores))}))

(comment
  
  (day2)

  (time (day2))
  


  )