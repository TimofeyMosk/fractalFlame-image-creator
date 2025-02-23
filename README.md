# Fractal Flame Image Creator

**Fractal Flame Image Creator** — консольное приложение для создания красивых абстрактных изображений с помощью
математиечских функций и СИФ.

## Возможности

- **Генерация изображений** заданной ширины и высоты
- **Логарифмическая гамма коррекция:** сглаживает цвета.
- **Stretching-Compression:** убирает лишние шумы на изображении.
- **Многопоточность:** Ускоряет вычисления и генерацию изображения.
- **Гибкость конфигурации:** поддержка различных параметров через командную строку.

## Установка

Убедитесь, что у вас установлен Go версии 1.22.6 или выше.

1. Склонируйте репозиторий:

    ```bash
    git clone ("current repo"*)
    cd "repo folder"*
    ```

\* - replace path

2. Соберите приложение:

    ```bash
    go build -o fractalFlameImageCreator ./cmd/run
    ```

3. Запустите с помощью:

    ```bash
    ./fractalFlameImageCreator [параметры]
    ```

## Использование

Приложение принимает следующие параметры:

| Параметр                  | Тип    | Описание                                                               | Обязательный |
|---------------------------|--------|------------------------------------------------------------------------|--------------|
| `-height`                 | int    | Высота изображения                                                     | ✅            |
| `-width`                  | int    | Ширина изображения                                                     | ✅            |
| `-iter`                   | int    | Количество итераций для генерации изображения                          | ✅            |
| `-linear-transform-count` | int    | Количество линейных преобразований                                     | ❌            |
| `-symmetry`               | -      | Добавляет симметрию                                                    | ❌            |
| `-threads`                | int    | Количество запускаемых горутин                                         | ❌            |
| `-scc`                    | int    | Коэффициент сжатия-растяжения изображения(убирает шумы на изображении) | ❌            |
| `-filename`               | string | Имя сохраняемого файла                                                 | ❌            |
| `-nonlinear-transforms`   | string | Имена нелинейных преобразований с их вероятностями                     | ❌            |
| `-log-gamma`              | -      | Добавляет логарифмическую гамма коррекцию                              | ❌            |
| `-gamma`                  | float  | Коэффициент gamma для log-gamma                                        | ❌            |

Значения по умолчанию
linear-transform-count = 10  
symmetry - false  
threads - Количество доступных логических ядер cpu  
scc - 1  
filename - "fractal_image_D_09_12_2024_T_01_42_59.png" (дата и время текущее)  
nonlinear-transforms - 4 Случайных трансформаций с случайными вероятностями  
-log-gamma - false  
gamma - 2.2 (сработает при наличии -log-gamma)  

## Список доступных преобразований

### **Name nonlinear transformation**: disk, handkerchief, heart, horseshoe, polar, sinusoidal, spherical, swirl, bubble, cosine, cross, diamond, exponential, eyefish, fisheye, hyperbolic, spiral, tangent

## Пример

```bash
./fractalFlameImageCreator -height=1080 -width=1920 -iter=50000000 -linear-transform-count=10 -symmetry -threads 8  -log-gamma -gamma 2.5 -nonlinear-transforms="sinusoidal:0.1,polar:0.1,disk:0.4,handkerchief: 0.4" -scc 5 -filename "Fractal.png"
```
```bash
./fractalFlameImageCreator  -height=1600 -width=2560 -iter=10000000   -log-gamma  scc=3
```
В этом примере:
Высота изображения - 1080,  
Ширина изображения - 1920,  
Количество итераций - 500000000,  
Количество линейных трансформаций - 10,  
Симметрия включена,  
Количество потоков(горутин) - 8,  
Логарифмичкая гамма коррекция включена, гамма - 2.5,  
Будут применены нелинейные преобразования:

- sinusoidal c вероятностью 10%,
- polar c вероятностью 10%,
- disk с вероятностью 10%,
- handkerchief с вероятностью 40%.

Имя файла - Fractal.png  
Коэффициент растяжения-сжатия - 5 (пикселей в процессе работы будет в 25 раз, на выходе изображение возвращается в
нужный формат)  

## Сравнение производительности

```
goos: darwin
goarch: arm64
pkg: github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/application
cpu: Apple M1
BenchmarkFractalFlameImageGenerator_Start
BenchmarkFractalFlameImageGenerator_Start/ThreadCount=1
BenchmarkFractalFlameImageGenerator_Start/ThreadCount=1-8         	       1	2306278417 ns/op
BenchmarkFractalFlameImageGenerator_Start/ThreadCount=2
BenchmarkFractalFlameImageGenerator_Start/ThreadCount=2-8         	       1	1036877583 ns/op
BenchmarkFractalFlameImageGenerator_Start/ThreadCount=4
BenchmarkFractalFlameImageGenerator_Start/ThreadCount=4-8         	       2	 575703604 ns/op
testing: BenchmarkFractalFlameImageGenerator_Start/ThreadCount=4-8 left GOMAXPROCS set to 4
BenchmarkFractalFlameImageGenerator_Start/ThreadCount=8
BenchmarkFractalFlameImageGenerator_Start/ThreadCount=8-8         	       3	 373330972 ns/op
BenchmarkFractalFlameImageGenerator_Start/ThreadCount=16
BenchmarkFractalFlameImageGenerator_Start/ThreadCount=16-8        	       3	 381136306 ns/op
```

goos: darwin: Система, на которой запущен тест (macOS).
goarch: arm64: Архитектура процессора (Apple M1 использует ARM64).
pkg: Пакет, в котором выполнялся тест.
cpu: Apple M1: Указание конкретного CPU.

При увеличении количества горутин производительность значительно растет до 8 потоков.
После 8 потоков эффективность перестает увеличиваться, так как Apple M1 имеет 8 производительных ядер,
и дополнительное увеличение потоков приводит к накладным расходам на управление. 
Тем не менее, было замечено, что иногда дополнительные горутины сверх числа физический ядер, 
даёт прирост производительности, но очень небольшой, и далеко не всегда(зависит от того, насколько сумма преобразований зациклена,
если сильно, то планировщик горутин находит ту, которая не заблокирована быстрее, уменьшая время прерывания).


Примеры изображений:
![fractal_image_D_09_12_2024_T_01_42_59.png](img/fractal_image_D_09_12_2024_T_01_42_59.png)
![fractal_image_D_10_12_2024_T_13_16_07.png](img/fractal_image_D_10_12_2024_T_13_16_07.png)
![fractal_image_D_10_12_2024_T_18_57_29.png](img/fractal_image_D_10_12_2024_T_18_57_29.png)
![fractal_image_D_10_12_2024_T_19_19_09.png](img/fractal_image_D_10_12_2024_T_19_19_09.png)
![fractal_image_D_10_12_2024_T_19_19_58.png](img/fractal_image_D_10_12_2024_T_19_19_58.png)
![fractal_image_D_10_12_2024_T_19_22_31.png](img/fractal_image_D_10_12_2024_T_19_22_31.png)
![fractal_image_D_10_12_2024_T_19_37_14.png](img/fractal_image_D_10_12_2024_T_19_37_14.png)
![fractal_image_D_11_12_2024_T_18_55_47.png](img/fractal_image_D_11_12_2024_T_18_55_47.png)