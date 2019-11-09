using System;

public class Program
{
	public static void Main()
	{
		double x = 1.0;
		double y = 1.0;
		Console.WriteLine("Enter X");
		x = double.Parse(Console.ReadLine());

		Console.WriteLine("Enter Y");
		y = double.Parse(Console.ReadLine());

		Console.WriteLine(isInTarget(x, y));
	}

	public static double parabola(double x)
	{
		return x * x;
	}

	public static double line(double x)
	{
		return 2 - x;
	}

	public static bool isInSectorD(double x, double y)
  {
		return x >= 0 && y >= 0 && isPositionUnderTheLine(x, y) && !isPositionInsideTheParabola(x, y);
	}

	public static bool isPositionUnderTheLine(double x, double y)
  {
		return line(x) >= y;
	}

	public static bool isPositionInsideTheParabola(double x, double y)
  {
		return !(parabola(x) > y);
	}

	public static bool isInSectorC(double x, double y)
  {
		return isPositionUnderTheLine(x, y) && isPositionInsideTheParabola(x, y);
	}

	public static bool isInTarget(double x, double y)
  {
		return isInSectorC(x, y) || isInSectorD(x, y);
	}
}
