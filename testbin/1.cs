// Hello World! program
namespace HelloWorld
{
    class Hello {         
        static void Main(string[] args)
        {
            System.Console.WriteLine("Hello World!");
            System.Console.WriteLine("Hello World111!");
            foreach (var arg in args)
                {
                    System.Console.WriteLine(arg);
                }
        }
    }
}