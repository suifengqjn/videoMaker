package app

import (
	"fmt"
	"math/rand"
	"myTool/ffmpeg"
	"myTool/file"
	"os"
	"strings"
)

var IsDealing bool
func (a *App) DoFactory() {

	if IsDealing {
		fmt.Println("软件正在处理中，请处理完再开始新的处理")
		return
	}

	videoDirs := a.getVideoDirs()

	if len(videoDirs) > 0 {
		fmt.Println("进行视频原创生成...")
		IsDealing = true
		fmt.Printf("本次读取到 %v 个文件夹\n",len(videoDirs))
		for i, f := range videoDirs {
			fmt.Printf("处理第 %v 个文件夹\n", i+1)
			a.doCompositeVideo(f)
		}
		fmt.Printf("视频原创处理结束, 一共处理了 %v 个文件夹\n\n\n", len(videoDirs))

	}

	dir := a.AppDir + "/video"
	videos := getVideos(dir)
	if len(videos) > 0 {
		IsDealing = true
		a.doEditVideo(videos, dir)
		fmt.Println("视频后期处理结束")
	}

	if len(videoDirs) == 0 && len(videos) == 0 {
		fmt.Println("没有视频需要处理")
	}

	IsDealing = false


}

func (a *App) doEditVideo(videos []string, dir string) {
	for _, f := range videos {
		a.postEdit(f, dir)
	}
}


func (a *App) doCompositeVideo(dir string) {

	//预处理
	//f = a.prepareEdit(f)

	//原创处理
	videoPath := ""
	if a.Composite.Switch {
		videoPath = a.composite(dir)
	} else {
		fmt.Println("本次没有选择原创选项")
	}

	if len(videoPath) > 0 {  //合成时候处理
		// 后处理
		videoPath = a.postEdit(videoPath, dir)
	}
}

func (a *App) prepareEdit(f string) string {
	fCmd := a.FCmd
	tempf := f
	if a.CutFront.Switch && a.CutBack.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			if a.CutFront.Value > 0 && a.CutBack.Value > 0 {
				f = info.CutFrontAndBack(fCmd, f, a.CutFront.Value, a.CutBack.Value)
			}
		}

	} else {
		//4. cut front
		if a.CutFront.Switch {
			info, err := ffmpeg.GetVideoInfo(fCmd, f)
			if err == nil && a.CutFront.Value > 0 {
				f = info.CutFront(fCmd, f, a.CutFront.Value)
			}

		}

		//5. cut back
		if a.CutBack.Switch {
			info, err := ffmpeg.GetVideoInfo(fCmd, f)
			if err == nil && a.CutBack.Value > 0 {
				f = info.CutBack(fCmd, f, a.CutBack.Value)
			}

		}
	}

	//去除水印
	if a.ClearWater.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			f = info.ClearWater(fCmd, f, a.ClearWater.X, a.ClearWater.Y, a.ClearWater.W, a.ClearWater.H)
		}

	}

	if a.ClearWater1.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			f = info.ClearWater(fCmd, f, a.ClearWater1.X, a.ClearWater1.Y, a.ClearWater1.W, a.ClearWater1.H)
		}

	}

	//覆盖文件
	if tempf != f {
		os.Remove(tempf)
		file.MoveFile(f, tempf)
	}
	return tempf
}


func (a *App)postEdit(videoPath, dir string) string  {

	if file.PathExist(videoPath) == false {
		return ""
	}
	fCmd := a.FCmd
	f := videoPath
	tempf := videoPath
	fmt.Println("进行视频后期处理",videoPath)
	if a.WaterText.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			alpha := a.WaterText.Alpha
			if alpha == 0 {
				alpha = 1.0
			}

			if a.WaterText.Content == "" {
				a.WaterText.Content = file.GetFileBaseName(f)
			}

			font := a.FontPath()
			if len(a.WaterText.Path) > 0 {
				font = a.WaterText.Path
			}

			f = info.AddTextWaterWithStyle(
				fCmd,
				f,
				font,
				a.WaterText.Size,
				a.WaterText.Content,
				a.WaterText.Style,
				a.WaterText.Sp1,
				a.WaterText.Sp2,
				a.WaterText.Color,
				alpha,
			)
		}

	}

	if a.RunWaterText.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			if a.RunWaterText.Content == "" {
				a.RunWaterText.Content = file.GetFileBaseName(f)
			}
			font := a.FontPath()
			if len(a.RunWaterText.Path) > 0 {
				font = a.RunWaterText.Path
			}
			f = info.AddScrollTextWater(
				fCmd,
				f,
				font,
				a.RunWaterText.Content,
				a.RunWaterText.Color,
				a.RunWaterText.Size,
				a.RunWaterText.IsTop == 2,
				a.RunWaterText.LeftToRight == 2,
				a.RunWaterText.Sp,
			)
		}

	}
	//10. water image
	if a.WaterImage.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			f = info.AddTextWaterImageWithStyle(
				fCmd,
				f,
				a.WaterImage.Path,
				a.WaterImage.Style,
				a.WaterImage.Sp1,
				a.WaterImage.Sp2,
			)
		}

	}


	// 添加背景音乐
	if a.AddBgm.Switch {
		f = a.addBgm(a.BgmDir(), f, false)
	}


	//11. film title
	if a.FilmHead.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			filmPath := a.AppDir + strings.TrimPrefix(a.FilmHead.Path, ".")
			if file.PathExist(filmPath) == false {
				fmt.Println("片头文件有误")
			} else {
				newHeader := ffmpeg.UpdateVideoSize(fCmd, filmPath, info.W, info.H)
				f = info.MergeVideoHeader(fCmd, newHeader, f)
			}

		}

	}
	//12. film end
	if a.FilmFoot.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			newFooter := ffmpeg.UpdateResolution(fCmd, a.FilmFoot.Path, info.W, info.H)
			if file.PathExist(newFooter) == false {
				fmt.Println("片尾文件有误")
			} else {
				f = info.MergeVideoFooter(fCmd, newFooter, f)
			}
		}

	}
	if tempf != f {
		result := fmt.Sprintf("%v/%v_result.%v", dir, file.GetFileBaseName(f),file.GetFileSuf(f))
		file.MoveFile(f, result)
	}
	return ""
}

func (a *App)addBgm(bgmPath ,videoPath string, cover bool)string  {

	bgmFiles := GetAllBgm(bgmPath)
	if len(bgmFiles) == 0 {
		fmt.Println("背景音乐为空")
		return videoPath
	}

	index := rand.Int() % len(bgmFiles)

	bgm :=  bgmFiles[index]

	info, err := ffmpeg.GetVideoInfo(a.FCmd,videoPath)
	if err != nil {
		return videoPath
	}

	return info.AddBgm(a.FCmd,videoPath,bgm,cover)

}


func GetAllBgm(dir string) []string  {
	files , err := file.GetAllFiles(dir)
	if err != nil {
		return nil
	}
	var res []string

	for _, f := range  files {
		v := ffmpeg.IsMusic(f)
		if v {
			res = append(res, f)
		}
	}
	return res
}