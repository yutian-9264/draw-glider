package main

import "github.com/fogleman/gg"

func main() {
    const W = 1024
    const H = 1024
    dc := gg.NewContext(W, H)
    dc.SetHexColor("fff")
    dc.Clear()

    dc.SetRGBA(0,0,0,1)
    dc.SetLineWidth(3)
    dc.DrawLine(W*0.5-150,H*0.5+150,W*0.5+150,H*0.5+150)
    dc.DrawLine(W*0.5-150,H*0.5+50,W*0.5+150,H*0.5+50)
    dc.DrawLine(W*0.5-150,H*0.5-50,W*0.5+150,H*0.5-50)
    dc.DrawLine(W*0.5-150,H*0.5-150,W*0.5+150,H*0.5-150)
    dc.DrawLine(W*0.5-150,H*0.5+150,W*0.5-150,H*0.5-150)
    dc.DrawLine(W*0.5-50,H*0.5+150,W*0.5-50,H*0.5-150)
    dc.DrawLine(W*0.5+50,H*0.5+150,W*0.5+50,H*0.5-150)
    dc.DrawLine(W*0.5+150,H*0.5+150,W*0.5+150,H*0.5-150)

    dc.Stroke()

    dc.SetRGBA(0,0,1,1)
    dc.DrawCircle(W*0.5,H*0.5-100, 40)
    dc.DrawCircle(W*0.5+100,H*0.5, 40)
    dc.DrawCircle(W*0.5-100,H*0.5+100, 40)
    dc.DrawCircle(W*0.5,H*0.5+100, 40)
    dc.DrawCircle(W*0.5+100,H*0.5+100, 40)
    dc.Fill()
    dc.Stroke()

    dc.SavePNG("out.png")
}
