using Microsoft.AspNetCore.Mvc;

namespace dotnet_old.Controllers;

[ApiController]
[Route("/")]
public class HelloController : ControllerBase
{

  public HelloController()
  {
  }

  [HttpGet]
  public HelloResponse Get()
  {
    return new HelloResponse("world");
  }
}
