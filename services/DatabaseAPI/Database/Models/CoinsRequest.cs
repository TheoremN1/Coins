using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace DatabaseAPI.Database.Models;

public class CoinsRequest
{
	[Key]
	[DatabaseGenerated(DatabaseGeneratedOption.Identity)]
	public int Id { get; set; }

	public int UserId { get; set; }

	public string UserMessage { get; set; }

	public int? HrId { get; set; }

	public string? HrMessage { get; set; }

	public int AchievementId { get; set; }

	public string StatusKey { get; set; }


	[ForeignKey(nameof(UserId))]
	public User? User { get; set; }

	[ForeignKey(nameof(HrId))]
	public User? Hr { get; set; }

	[ForeignKey(nameof(AchievementId))]
	public Achievement? Achievement { get; set; }

	[ForeignKey(nameof(StatusKey))]
	public Status? Status { get; set; }
}
